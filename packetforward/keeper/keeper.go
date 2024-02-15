package keeper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/armon/go-metrics"
	"github.com/sideprotocol/packet-forward-middleware/v7/packetforward/types"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	"github.com/cometbft/cometbft/libs/log"

	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
)

var (
	// DefaultTransferPacketTimeoutHeight is the timeout height following IBC defaults
	DefaultTransferPacketTimeoutHeight = clienttypes.Height{
		RevisionNumber: 0,
		RevisionHeight: 0,
	}

	// DefaultForwardTransferPacketTimeoutTimestamp is the timeout timestamp following IBC defaults
	DefaultForwardTransferPacketTimeoutTimestamp = time.Duration(transfertypes.DefaultRelativePacketTimeoutTimestamp) * time.Nanosecond

	// DefaultRefundTransferPacketTimeoutTimestamp is a 28-day timeout for refund packets since funds are stuck in packetforward module otherwise.
	DefaultRefundTransferPacketTimeoutTimestamp = 28 * 24 * time.Hour
)

// Keeper defines the packet forward middleware keeper
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey

	transferKeeper types.TransferKeeper
	channelKeeper  types.ChannelKeeper
	distrKeeper    types.DistributionKeeper
	bankKeeper     types.BankKeeper
	ics4Wrapper    porttypes.ICS4Wrapper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new forward Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	transferKeeper types.TransferKeeper,
	channelKeeper types.ChannelKeeper,
	distrKeeper types.DistributionKeeper,
	bankKeeper types.BankKeeper,
	ics4Wrapper porttypes.ICS4Wrapper,
	authority string,
) *Keeper {
	return &Keeper{
		cdc:            cdc,
		storeKey:       key,
		transferKeeper: transferKeeper,
		channelKeeper:  channelKeeper,
		distrKeeper:    distrKeeper,
		bankKeeper:     bankKeeper,
		ics4Wrapper:    ics4Wrapper,
		authority:      authority,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetTransferKeeper sets the transferKeeper
func (k *Keeper) SetTransferKeeper(transferKeeper types.TransferKeeper) {
	k.transferKeeper = transferKeeper
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+ibcexported.ModuleName+"-"+types.ModuleName)
}

func (k *Keeper) WriteAcknowledgementForForwardedPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data types.ICS101SwapMsg,
	multihopsPacket *types.MultiHopsPacket,
	ack channeltypes.Acknowledgement,
) error {
	// Lookup module by channel capability
	_, chanCap, err := k.channelKeeper.LookupModuleByChannel(ctx, multihopsPacket.PacketSrcPortId, multihopsPacket.PacketSrcChannelId)
	if err != nil {
		return errorsmod.Wrap(err, "could not retrieve module from port-id")
	}

	// for forwarded packets, the funds were moved into an escrow account if the denom originated on this chain.
	// On an ack error or timeout on a forwarded packet, the funds in the escrow account
	// should be moved to the other escrow account on the other side or burned.
	if !ack.Success() {

		ackResult := fmt.Sprintf("packet forward failed after point of no return: %s", ack.GetError())
		newAck := channeltypes.NewResultAcknowledgement([]byte(ackResult))

		return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, channeltypes.Packet{
			Data:               multihopsPacket.PacketData,
			Sequence:           multihopsPacket.PacketSequence,
			SourcePort:         multihopsPacket.PacketSrcPortId,
			SourceChannel:      multihopsPacket.PacketSrcChannelId,
			DestinationPort:    multihopsPacket.PacketDestPortId,
			DestinationChannel: multihopsPacket.PacketDestChannelId,
			TimeoutHeight:      clienttypes.MustParseHeight(multihopsPacket.PacketTimeoutHeight),
			TimeoutTimestamp:   multihopsPacket.PacketTimeoutTimestamp,
		}, newAck)
	}

	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, channeltypes.Packet{
		Data:               multihopsPacket.PacketData,
		Sequence:           multihopsPacket.PacketSequence,
		SourcePort:         multihopsPacket.PacketSrcPortId,
		SourceChannel:      multihopsPacket.PacketSrcChannelId,
		DestinationPort:    multihopsPacket.PacketDestPortId,
		DestinationChannel: multihopsPacket.PacketDestChannelId,
		TimeoutHeight:      clienttypes.MustParseHeight(multihopsPacket.PacketTimeoutHeight),
		TimeoutTimestamp:   multihopsPacket.PacketTimeoutTimestamp,
	}, ack)
}

func (k *Keeper) ForwardPacket(
	ctx sdk.Context,
	multihopsPacket *types.MultiHopsPacket,
	srcPacket channeltypes.Packet,
	data types.ICS101SwapMsg,
	metadata *types.ForwardMetadata,
	maxRetries uint8,
	timeout time.Duration,
	labels []metrics.Label,
) error {
	var err error
	memo := ""
	// set memo for next transfer with next from this transfer.
	if metadata.Next != nil {
		memoBz, err := json.Marshal(metadata.Next)
		if err != nil {
			k.Logger(ctx).Error("packetForwardMiddleware error marshaling next as JSON",
				"error", err,
			)
			return errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
		}
		memo = string(memoBz)
	}

	_, chanCap, err := k.channelKeeper.LookupModuleByChannel(ctx, metadata.Port, metadata.Channel)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}
	seq, err := k.SendPacket(ctx, chanCap, metadata.Port, metadata.Channel, DefaultTransferPacketTimeoutHeight, uint64(ctx.BlockTime().UnixNano())+uint64(timeout.Nanoseconds()), srcPacket.Data)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}

	// Update memo
	data.Memo = memo
	newData, err := json.Marshal(data)
	srcPacket.Data = newData

	k.Logger(ctx).Debug("packetForwardMiddleware ForwardTransferPacket",
		"port", metadata.Port, "channel", metadata.Channel,
		"receiver", metadata.Receiver,
	)

	if err != nil {
		k.Logger(ctx).Error("packetForwardMiddleware ForwardTransferPacket error",
			"port", metadata.Port, "channel", metadata.Channel,
			"receiver", metadata.Receiver,
			"error", err,
		)
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	// Store the following information in keeper:
	// key - information about forwarded packet: src_channel (parsedReceiver.Channel), src_port (parsedReceiver.Port), sequence
	// value - information about original packet for refunding if necessary: retries, srcPacketSender, srcPacket.DestinationChannel, srcPacket.DestinationPort

	if multihopsPacket == nil {
		multihopsPacket = &types.MultiHopsPacket{
			PacketData:             srcPacket.Data,
			PacketSrcPortId:        srcPacket.SourcePort,
			PacketSrcChannelId:     srcPacket.SourceChannel,
			PacketTimeoutTimestamp: srcPacket.TimeoutTimestamp,
			PacketTimeoutHeight:    srcPacket.TimeoutHeight.String(),
			RetriesRemaining:       int32(maxRetries),
			Timeout:                uint64(timeout.Nanoseconds()),
		}
	} else {
		multihopsPacket.RetriesRemaining--
	}

	key := types.RefundPacketKey(metadata.Channel, metadata.Port, seq)
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(multihopsPacket)
	store.Set(key, bz)

	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"ibc", types.ModuleName, "send"},
			1,
			labels,
		)
	}()
	return nil
}

// TimeoutShouldRetry returns inFlightPacket and no error if retry should be attempted. Error is returned if IBC refund should occur.
func (k *Keeper) TimeoutShouldRetry(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (*types.MultiHopsPacket, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.RefundPacketKey(packet.SourceChannel, packet.SourcePort, packet.Sequence)

	if !store.Has(key) {
		// not a forwarded packet, ignore.
		return nil, nil
	}

	bz := store.Get(key)
	var multihopsPacket types.MultiHopsPacket
	k.cdc.MustUnmarshal(bz, &multihopsPacket)

	if multihopsPacket.RetriesRemaining <= 0 {
		k.Logger(ctx).Error("packetForwardMiddleware reached max retries for packet",
			"key", string(key),
			"original-sender-address", multihopsPacket.OriginalSenderAddress,
			"dest-channel-id", multihopsPacket.PacketDestChannelId,
			"dest-port-id", multihopsPacket.PacketDestPortId,
		)
		return &multihopsPacket, fmt.Errorf("giving up on packet on channel (%s) port (%s) after max retries",
			multihopsPacket.PacketDestChannelId, multihopsPacket.PacketDestPortId)
	}

	return &multihopsPacket, nil
}

func (k *Keeper) RetryTimeout(
	ctx sdk.Context,
	channel, port string,
	data types.ICS101SwapMsg,
	multihopsPacket *types.MultiHopsPacket,
) error {
	// send transfer again
	metadata := &types.ForwardMetadata{
		Channel: channel,
		Port:    port,
	}

	if data.Memo != "" {
		metadata.Next = &types.JSONObject{}
		if err := json.Unmarshal([]byte(data.Memo), metadata.Next); err != nil {
			return fmt.Errorf("error unmarshaling memo json: %w", err)
		}
	}
	// srcPacket and srcPacketSender are empty because inFlightPacket is non-nil.
	return k.ForwardPacket(
		ctx,
		multihopsPacket,
		channeltypes.Packet{},
		data,
		metadata,
		uint8(multihopsPacket.RetriesRemaining),
		time.Duration(multihopsPacket.Timeout)*time.Nanosecond,
		[]metrics.Label{},
	)
}

func (k *Keeper) RemoveMultiHopsPacket(ctx sdk.Context, packet channeltypes.Packet) {
	store := ctx.KVStore(k.storeKey)
	key := types.RefundPacketKey(packet.SourceChannel, packet.SourcePort, packet.Sequence)
	if !store.Has(key) {
		// not a forwarded packet, ignore.
		return
	}

	// done with packet key now, delete.
	store.Delete(key)
}

// GetAndClearInFlightPacket will fetch an InFlightPacket from the store, remove it if it exists, and return it.
func (k *Keeper) GetAndClearMultiHopsPacket(
	ctx sdk.Context,
	channel string,
	port string,
	sequence uint64,
) *types.MultiHopsPacket {
	store := ctx.KVStore(k.storeKey)
	key := types.RefundPacketKey(channel, port, sequence)
	if !store.Has(key) {
		// this is either not a forwarded packet, or it is the final destination for the refund.
		return nil
	}

	bz := store.Get(key)

	// done with packet key now, delete.
	store.Delete(key)

	var multihopsPacket types.MultiHopsPacket
	k.cdc.MustUnmarshal(bz, &multihopsPacket)
	return &multihopsPacket
}

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (k Keeper) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string, sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	return k.ics4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// WriteAcknowledgement wraps IBC ICS4Wrapper WriteAcknowledgement function.
// ICS29 WriteAcknowledgement is used for asynchronous acknowledgements.
func (k *Keeper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI, acknowledgement ibcexported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// WriteAcknowledgement wraps IBC ICS4Wrapper GetAppVersion function.
func (k *Keeper) GetAppVersion(
	ctx sdk.Context,
	portID,
	channelID string,
) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}

// LookupModuleByChannel wraps ChannelKeeper LookupModuleByChannel function.
func (k *Keeper) LookupModuleByChannel(ctx sdk.Context, portID, channelID string) (string, *capabilitytypes.Capability, error) {
	return k.channelKeeper.LookupModuleByChannel(ctx, portID, channelID)
}
