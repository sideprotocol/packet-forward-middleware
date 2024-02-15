package types

// NewGenesisState creates a pfm GenesisState instance.
func NewGenesisState(params Params, multiHopsPacket map[string]MultiHopsPacket) *GenesisState {
	return &GenesisState{
		Params:           params,
		MultiHopsPackets: multiHopsPacket,
	}
}

// DefaultGenesisState returns a GenesisState with a default fee percentage of 0.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		MultiHopsPackets: make(map[string]MultiHopsPacket),
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
