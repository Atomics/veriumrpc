package veriumrpc

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/url"
)

type WalletConf struct {
	Host     string
	Port     int64
	User     string
	Password string
}

type Wallet struct {
	httpClient *http.Client
	conf       *WalletConf
}

func New(walletConf *WalletConf) (*Wallet, error) {

	httpClient, err := newHTTPClient(walletConf)
	if err != nil {
		return nil, err
	}

	wallet := &Wallet{
		httpClient: httpClient,
		conf:       walletConf,
	}

	return wallet, nil
}

// Nothing usefull right now since i don't manage proxy and tls is not supported by verium
func newHTTPClient(walletConf *WalletConf) (*http.Client, error) {
	var proxyFunc func(*http.Request) (*url.URL, error)
	var tlsConfig *tls.Config

	client := http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
		},
	}

	return &client, nil
}

func (wallet *Wallet) sendPost(jReq *jsonRequest) {

	protocol := "http"
	url := protocol + "://" + wallet.conf.Host
	bodyReader := bytes.NewReader(jReq.marshalledJSON)
	httpReq, err := http.NewRequest("POST", url, bodyReader)
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	// Configure basic access authorization.
	httpReq.SetBasicAuth(wallet.conf.User, wallet.conf.Password)

	// To Continue ....
}
