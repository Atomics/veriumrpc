package veriumrpc

type Block struct {
	Hash              string
	Confirmation      int64
	Size              int64
	Height            int64
	Version           int64
	Merkleroot        string
	Mint              float64
	Time              int64
	Nonce             int64
	Bits              string
	Difficulty        float64
	Blocktrust        string
	Chaintrust        string
	Previousblockhash string
	Nextblockhash     string
	Flags             string
	Proofhash         string
}
