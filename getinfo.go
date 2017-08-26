package veriumrpc

import (
	"encoding/json"
)

type getInfo struct {
	Version         string  `json:"version"`
	ProtocolVersion int64   `json:"protocolversion"`
	WalletVersion   int64   `json:"walletversion"`
	Balance         float64 `json:"balance"`
	NewMint         float64 `json:"newmint"`
	TotalSupply     float64 `json:"totalsupply"`
	Blocks          int64   `json:"blocks"`
	Timeoffset      int64   `json:"timeoffset"`
	Connections     int64   `json:"connections"`
	Proxy           string  `json:"proxy"`
	Ip              string  `json:"ip"`
	Difficulty      float64 `json:"difficulty"`
	BlocksPerHour   int64   `json:"blocksperhour"`
	Testnet         bool    `json:"testnet"`
	KeyPoolOldest   int64   `json:"keypoololdest"`
	KeyPoolSize     int64   `json:"keypoolsize"`
	PayTxFee        float64 `json:"paytxfee"`
	MinInput        float64 `json:"mininput"`
	Errors          string  `json:"errors"`
}

func (wallet *Wallet) GetInfo() (*getInfo, error) {
	result, err := wallet.sendPost(
		&jsonRequest{
			Jsonrpc: "1.0",
			Id:      "getinfo",
			Method:  "getinfo",
			Params:  make([]interface{}, 0),
		},
	)
	if err != nil {
		return nil, err
	}

	var getInfoData getInfo
	err = json.Unmarshal(*result, &getInfoData)
	if err != nil {
		return nil, err
	}

	return &getInfoData, nil
}
