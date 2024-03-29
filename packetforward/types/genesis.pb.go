// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: packetforward/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the packetforward genesis state
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// key - information about forwarded packet: src_channel
	// (parsedReceiver.Channel), src_port (parsedReceiver.Port), sequence value -
	// information about original packet for refunding if necessary: retries,
	// srcPacketSender, srcPacket.DestinationChannel, srcPacket.DestinationPort
	MultiHopsPackets map[string]MultiHopsPacket `protobuf:"bytes,2,rep,name=multi_hops_packets,json=multiHopsPackets,proto3" json:"multi_hops_packets" yaml:"multi_hops_packets" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd4e56ea31af982, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMultiHopsPackets() map[string]MultiHopsPacket {
	if m != nil {
		return m.MultiHopsPackets
	}
	return nil
}

// Params defines the set of packetforward parameters.
type Params struct {
	FeePercentage github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=fee_percentage,json=feePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_percentage" yaml:"fee_percentage"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd4e56ea31af982, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// InFlightPacket contains information about original packet for
// writing the acknowledgement and refunding if necessary.
type MultiHopsPacket struct {
	OriginalSenderAddress  string `protobuf:"bytes,1,opt,name=original_sender_address,json=originalSenderAddress,proto3" json:"original_sender_address,omitempty"`
	PacketDestChannelId    string `protobuf:"bytes,2,opt,name=packet_dest_channel_id,json=packetDestChannelId,proto3" json:"packet_dest_channel_id,omitempty"`
	PacketDestPortId       string `protobuf:"bytes,3,opt,name=packet_dest_port_id,json=packetDestPortId,proto3" json:"packet_dest_port_id,omitempty"`
	PacketSrcChannelId     string `protobuf:"bytes,4,opt,name=packet_src_channel_id,json=packetSrcChannelId,proto3" json:"packet_src_channel_id,omitempty"`
	PacketSrcPortId        string `protobuf:"bytes,5,opt,name=packet_src_port_id,json=packetSrcPortId,proto3" json:"packet_src_port_id,omitempty"`
	PacketTimeoutTimestamp uint64 `protobuf:"varint,6,opt,name=packet_timeout_timestamp,json=packetTimeoutTimestamp,proto3" json:"packet_timeout_timestamp,omitempty"`
	PacketTimeoutHeight    string `protobuf:"bytes,7,opt,name=packet_timeout_height,json=packetTimeoutHeight,proto3" json:"packet_timeout_height,omitempty"`
	PacketData             []byte `protobuf:"bytes,8,opt,name=packet_data,json=packetData,proto3" json:"packet_data,omitempty"`
	PacketSequence         uint64 `protobuf:"varint,9,opt,name=packet_sequence,json=packetSequence,proto3" json:"packet_sequence,omitempty"`
	RetriesRemaining       int32  `protobuf:"varint,10,opt,name=retries_remaining,json=retriesRemaining,proto3" json:"retries_remaining,omitempty"`
	Timeout                uint64 `protobuf:"varint,11,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (m *MultiHopsPacket) Reset()         { *m = MultiHopsPacket{} }
func (m *MultiHopsPacket) String() string { return proto.CompactTextString(m) }
func (*MultiHopsPacket) ProtoMessage()    {}
func (*MultiHopsPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_afd4e56ea31af982, []int{2}
}
func (m *MultiHopsPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiHopsPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiHopsPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiHopsPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiHopsPacket.Merge(m, src)
}
func (m *MultiHopsPacket) XXX_Size() int {
	return m.Size()
}
func (m *MultiHopsPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiHopsPacket.DiscardUnknown(m)
}

var xxx_messageInfo_MultiHopsPacket proto.InternalMessageInfo

func (m *MultiHopsPacket) GetOriginalSenderAddress() string {
	if m != nil {
		return m.OriginalSenderAddress
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketDestChannelId() string {
	if m != nil {
		return m.PacketDestChannelId
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketDestPortId() string {
	if m != nil {
		return m.PacketDestPortId
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketSrcChannelId() string {
	if m != nil {
		return m.PacketSrcChannelId
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketSrcPortId() string {
	if m != nil {
		return m.PacketSrcPortId
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketTimeoutTimestamp() uint64 {
	if m != nil {
		return m.PacketTimeoutTimestamp
	}
	return 0
}

func (m *MultiHopsPacket) GetPacketTimeoutHeight() string {
	if m != nil {
		return m.PacketTimeoutHeight
	}
	return ""
}

func (m *MultiHopsPacket) GetPacketData() []byte {
	if m != nil {
		return m.PacketData
	}
	return nil
}

func (m *MultiHopsPacket) GetPacketSequence() uint64 {
	if m != nil {
		return m.PacketSequence
	}
	return 0
}

func (m *MultiHopsPacket) GetRetriesRemaining() int32 {
	if m != nil {
		return m.RetriesRemaining
	}
	return 0
}

func (m *MultiHopsPacket) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "packetforward.v1.GenesisState")
	proto.RegisterMapType((map[string]MultiHopsPacket)(nil), "packetforward.v1.GenesisState.MultiHopsPacketsEntry")
	proto.RegisterType((*Params)(nil), "packetforward.v1.Params")
	proto.RegisterType((*MultiHopsPacket)(nil), "packetforward.v1.MultiHopsPacket")
}

func init() { proto.RegisterFile("packetforward/v1/genesis.proto", fileDescriptor_afd4e56ea31af982) }

var fileDescriptor_afd4e56ea31af982 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcf, 0x4e, 0xdb, 0x30,
	0x1c, 0x6e, 0x4a, 0x29, 0xc3, 0x65, 0xd0, 0x99, 0x75, 0xcb, 0x38, 0xa4, 0xa5, 0x87, 0xad, 0x12,
	0x6a, 0x22, 0x60, 0x02, 0xc4, 0x6d, 0x1d, 0x13, 0x70, 0x98, 0x54, 0x05, 0x4e, 0xbb, 0x44, 0x26,
	0xf9, 0x35, 0x8d, 0x48, 0xe2, 0xcc, 0x76, 0xcb, 0xfa, 0x00, 0xbb, 0xed, 0xb0, 0xa7, 0xd9, 0x33,
	0x70, 0xe4, 0x38, 0xed, 0x50, 0x4d, 0xf0, 0x06, 0x3c, 0xc1, 0x54, 0xdb, 0x81, 0xd2, 0xee, 0x14,
	0xdb, 0xdf, 0x9f, 0xdf, 0xf7, 0x59, 0x6e, 0x91, 0x95, 0x11, 0xff, 0x12, 0x44, 0x8f, 0xb2, 0x2b,
	0xc2, 0x02, 0x67, 0xb8, 0xed, 0x84, 0x90, 0x02, 0x8f, 0xb8, 0x9d, 0x31, 0x2a, 0x28, 0xae, 0x3e,
	0xc1, 0xed, 0xe1, 0xf6, 0xc6, 0xcb, 0x90, 0x86, 0x54, 0x82, 0xce, 0x64, 0xa5, 0x78, 0xcd, 0x5f,
	0x45, 0xb4, 0x72, 0xac, 0x94, 0x67, 0x82, 0x08, 0xc0, 0x7b, 0xa8, 0x9c, 0x11, 0x46, 0x12, 0x6e,
	0x1a, 0x0d, 0xa3, 0x55, 0xd9, 0x31, 0xed, 0x59, 0x27, 0xbb, 0x2b, 0xf1, 0x4e, 0xe9, 0x7a, 0x5c,
	0x2f, 0xb8, 0x9a, 0x8d, 0xbf, 0x1b, 0x08, 0x27, 0x83, 0x58, 0x44, 0x5e, 0x9f, 0x66, 0xdc, 0x53,
	0x22, 0x6e, 0x16, 0x1b, 0x0b, 0xad, 0xca, 0xce, 0xfb, 0x79, 0x93, 0xe9, 0xa1, 0xf6, 0xe7, 0x89,
	0xf0, 0x84, 0x66, 0xbc, 0xab, 0x64, 0x9f, 0x52, 0xc1, 0x46, 0x9d, 0xcd, 0xc9, 0x80, 0xfb, 0x71,
	0xfd, 0xcd, 0x88, 0x24, 0xf1, 0x61, 0x73, 0xde, 0xbd, 0xe9, 0x56, 0x93, 0x19, 0xe5, 0x46, 0x0f,
	0xd5, 0xfe, 0xeb, 0x86, 0xab, 0x68, 0xe1, 0x12, 0x46, 0xb2, 0xd5, 0xb2, 0x3b, 0x59, 0xe2, 0x7d,
	0xb4, 0x38, 0x24, 0xf1, 0x00, 0xcc, 0xa2, 0x6c, 0xba, 0x39, 0x1f, 0x72, 0xc6, 0xc9, 0x55, 0xfc,
	0xc3, 0xe2, 0x81, 0xd1, 0xfc, 0x86, 0xca, 0xea, 0x1e, 0x70, 0x8a, 0x56, 0x7b, 0x00, 0x5e, 0x06,
	0xcc, 0x87, 0x54, 0x90, 0x10, 0xd4, 0x8c, 0xce, 0xf1, 0x24, 0xfe, 0x9f, 0x71, 0xfd, 0x6d, 0x18,
	0x89, 0xfe, 0xe0, 0xc2, 0xf6, 0x69, 0xe2, 0xf8, 0x94, 0x27, 0x94, 0xeb, 0x4f, 0x9b, 0x07, 0x97,
	0x8e, 0x18, 0x65, 0xc0, 0xed, 0x23, 0xf0, 0xef, 0xc7, 0xf5, 0x9a, 0x2a, 0xfa, 0xd4, 0xad, 0xe9,
	0x3e, 0xef, 0x01, 0x74, 0x1f, 0xf7, 0x3f, 0x4a, 0x68, 0x6d, 0x26, 0x18, 0xde, 0x43, 0xaf, 0x29,
	0x8b, 0xc2, 0x28, 0x25, 0xb1, 0xc7, 0x21, 0x0d, 0x80, 0x79, 0x24, 0x08, 0x18, 0x70, 0xae, 0x0b,
	0xd7, 0x72, 0xf8, 0x4c, 0xa2, 0x1f, 0x14, 0x88, 0x77, 0xd1, 0x2b, 0x55, 0xda, 0x0b, 0x80, 0x0b,
	0xcf, 0xef, 0x93, 0x34, 0x85, 0xd8, 0x8b, 0x02, 0x79, 0x27, 0xcb, 0xee, 0xba, 0x42, 0x8f, 0x80,
	0x8b, 0x8f, 0x0a, 0x3b, 0x0d, 0x70, 0x1b, 0xad, 0x4f, 0x8b, 0x32, 0xca, 0xc4, 0x44, 0xb1, 0x20,
	0x15, 0xd5, 0x47, 0x45, 0x97, 0x32, 0x71, 0x1a, 0xe0, 0x6d, 0x54, 0xd3, 0x74, 0xce, 0xfc, 0xe9,
	0x11, 0x25, 0x29, 0xc0, 0x0a, 0x3c, 0x63, 0xfe, 0xe3, 0x84, 0x2d, 0x84, 0xa7, 0x24, 0xf9, 0x80,
	0x45, 0xc9, 0x5f, 0x7b, 0xe0, 0x6b, 0xff, 0x03, 0x64, 0x6a, 0xb2, 0x88, 0x12, 0xa0, 0x03, 0xf5,
	0xe5, 0x82, 0x24, 0x99, 0x59, 0x6e, 0x18, 0xad, 0x92, 0xab, 0x3b, 0x9e, 0x2b, 0xf8, 0x3c, 0x47,
	0xf1, 0xce, 0x43, 0xb2, 0x5c, 0xd9, 0x87, 0x28, 0xec, 0x0b, 0x73, 0x69, 0xba, 0xbc, 0x96, 0x9d,
	0x48, 0x08, 0xd7, 0x51, 0x25, 0x2f, 0x4f, 0x04, 0x31, 0x9f, 0x35, 0x8c, 0xd6, 0x8a, 0x8b, 0x74,
	0x69, 0x22, 0x08, 0x7e, 0x87, 0xd6, 0xf2, 0xec, 0xf0, 0x75, 0x00, 0xa9, 0x0f, 0xe6, 0xb2, 0x4c,
	0xb1, 0xaa, 0x83, 0xeb, 0x53, 0xbc, 0x85, 0x5e, 0x30, 0x10, 0x2c, 0x02, 0xee, 0x31, 0x48, 0x48,
	0x94, 0x46, 0x69, 0x68, 0xa2, 0x86, 0xd1, 0x5a, 0x74, 0xab, 0x1a, 0x70, 0xf3, 0x73, 0x6c, 0xa2,
	0x25, 0x9d, 0xd1, 0xac, 0x48, 0xb7, 0x7c, 0xdb, 0xb9, 0xb8, 0xbe, 0xb5, 0x8c, 0x9b, 0x5b, 0xcb,
	0xf8, 0x7b, 0x6b, 0x19, 0x3f, 0xef, 0xac, 0xc2, 0xcd, 0x9d, 0x55, 0xf8, 0x7d, 0x67, 0x15, 0xbe,
	0x9c, 0x4c, 0x3d, 0x3c, 0x1e, 0x05, 0x20, 0x7f, 0xf1, 0x3e, 0x8d, 0x1d, 0x15, 0xa4, 0xad, 0x1f,
	0x7a, 0x3b, 0x89, 0x82, 0x20, 0x86, 0x2b, 0xc2, 0xc0, 0x19, 0xee, 0x3b, 0x4f, 0xff, 0x58, 0xe4,
	0xf3, 0xbc, 0x28, 0x4b, 0xe9, 0xee, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x04, 0xd0, 0xba,
	0x76, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MultiHopsPackets) > 0 {
		for k := range m.MultiHopsPackets {
			v := m.MultiHopsPackets[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeePercentage.Size()
		i -= size
		if _, err := m.FeePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MultiHopsPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiHopsPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MultiHopsPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timeout != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x58
	}
	if m.RetriesRemaining != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RetriesRemaining))
		i--
		dAtA[i] = 0x50
	}
	if m.PacketSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PacketSequence))
		i--
		dAtA[i] = 0x48
	}
	if len(m.PacketData) > 0 {
		i -= len(m.PacketData)
		copy(dAtA[i:], m.PacketData)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketData)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PacketTimeoutHeight) > 0 {
		i -= len(m.PacketTimeoutHeight)
		copy(dAtA[i:], m.PacketTimeoutHeight)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketTimeoutHeight)))
		i--
		dAtA[i] = 0x3a
	}
	if m.PacketTimeoutTimestamp != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PacketTimeoutTimestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PacketSrcPortId) > 0 {
		i -= len(m.PacketSrcPortId)
		copy(dAtA[i:], m.PacketSrcPortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketSrcPortId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PacketSrcChannelId) > 0 {
		i -= len(m.PacketSrcChannelId)
		copy(dAtA[i:], m.PacketSrcChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketSrcChannelId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PacketDestPortId) > 0 {
		i -= len(m.PacketDestPortId)
		copy(dAtA[i:], m.PacketDestPortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketDestPortId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PacketDestChannelId) > 0 {
		i -= len(m.PacketDestChannelId)
		copy(dAtA[i:], m.PacketDestChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PacketDestChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginalSenderAddress) > 0 {
		i -= len(m.OriginalSenderAddress)
		copy(dAtA[i:], m.OriginalSenderAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OriginalSenderAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.MultiHopsPackets) > 0 {
		for k, v := range m.MultiHopsPackets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + l + sovGenesis(uint64(l))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FeePercentage.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *MultiHopsPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginalSenderAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketDestChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketDestPortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketSrcChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketSrcPortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PacketTimeoutTimestamp != 0 {
		n += 1 + sovGenesis(uint64(m.PacketTimeoutTimestamp))
	}
	l = len(m.PacketTimeoutHeight)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PacketData)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.PacketSequence != 0 {
		n += 1 + sovGenesis(uint64(m.PacketSequence))
	}
	if m.RetriesRemaining != 0 {
		n += 1 + sovGenesis(uint64(m.RetriesRemaining))
	}
	if m.Timeout != 0 {
		n += 1 + sovGenesis(uint64(m.Timeout))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiHopsPackets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MultiHopsPackets == nil {
				m.MultiHopsPackets = make(map[string]MultiHopsPacket)
			}
			var mapkey string
			mapvalue := &MultiHopsPacket{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &MultiHopsPacket{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MultiHopsPackets[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MultiHopsPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MultiHopsPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiHopsPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalSenderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalSenderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketDestChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketDestChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketDestPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketDestPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSrcChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketSrcChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSrcPortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketSrcPortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutTimestamp", wireType)
			}
			m.PacketTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketTimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketTimeoutHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketData = append(m.PacketData[:0], dAtA[iNdEx:postIndex]...)
			if m.PacketData == nil {
				m.PacketData = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketSequence", wireType)
			}
			m.PacketSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetriesRemaining", wireType)
			}
			m.RetriesRemaining = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetriesRemaining |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
