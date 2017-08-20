package veriumrpc

type getInfo struct {
	Version         string
	ProtocolVersion int64
	WalletVersion   int64
	Balance         float64
	NewMint         float64
	TotalSupply     float64
	Blocks          int64
	Timeoffest      int64
	Connections     int64
	Proxy           string
	Ip              string
	Difficulty      float64
	BlocksPerHour   int64
	Testnet         bool
	KeyPoolOldest   int64
	KeyPoolSize     int64
	PayTxFee        float64
	MinInput        float64
	Errors          string
}
