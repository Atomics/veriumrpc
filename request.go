package veriumrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type jsonRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type jsonResponse struct {
	Result *json.RawMessage `json:"result"`
	Error  string           `json:"error"`
	Id     string           `json:"id"`
}

func (wallet *Wallet) sendPost(requestData *jsonRequest) (*json.RawMessage, error) {
	url := fmt.Sprintf("http://%s:%d", wallet.conf.Host, wallet.conf.Port)

	jsonData, err := json.Marshal(&requestData)
	if err != nil {
		return nil, fmt.Errorf("Impossible to transform the requestData in json, Error: %v", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("Impossible to prepare the HTTP Request, Error: %v", err)
	}

	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(wallet.conf.User, wallet.conf.Password)

	httpResp, err := wallet.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Error during the HTTP Request, Error: %v", err)
	}
	defer httpResp.Body.Close()

	respBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("Impossible to read the json reply, Error: %v", err)
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		if len(respBytes) == 0 {
			return nil, fmt.Errorf("Wallet return a error code without data. Status Code: %d, Detail: %s",
				httpResp.StatusCode,
				http.StatusText(httpResp.StatusCode),
			)
		}
		return nil, fmt.Errorf("%s", respBytes)
	}

	var jsonResponseData jsonResponse
	err = json.Unmarshal(respBytes, &jsonResponseData)
	if err != nil {
		return nil, err
	}

	if jsonResponseData.Error != "" {
		return nil, errors.New(jsonResponseData.Error)
	}

	return jsonResponseData.Result, nil
}
