package types

// InterchainMessageType as an equivalent of Rust enum in Go
type InterchainMessageType int

const (
	Unspecified InterchainMessageType = iota
	MakePool
	TakePool
	CancelPool
	SingleAssetDeposit
	MakeMultiDeposit
	CancelMultiDeposit
	TakeMultiDeposit
	MultiWithdraw
	LeftSwap
	RightSwap
)

// InterchainSwapPacketData struct converted from Rust
type InterchainSwapPacketData struct {
	Type        InterchainMessageType `json:"Type"`
	Data        []byte                `json:"Data"`                  // Assuming Binary is a byte slice
	StateChange *[]byte               `json:"StateChange,omitempty"` // Pointer to slice for optional presence
	Memo        []byte                `json:"Memo,omitempty"`        // Pointer to slice for optional presence
}
