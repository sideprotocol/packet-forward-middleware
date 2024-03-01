package types

type WasmInterchainSwapPacketData struct {
	Type        string `json:"Type"`
	Data        string `json:"Data"`        // Base64 encoded string
	StateChange string `json:"StateChange"` // Base64 encoded string
	Memo        string `json:"Memo"`        // Base64 encoded string, optional
}
