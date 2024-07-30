package ztm_go_sdk

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// NewRestClient creates a Rest Client
func NewRestClient(serverAddr string, debug bool) *RestClient {
	return NewRestClientWithTransport(
		serverAddr,
		&http.Transport{
			DisableKeepAlives:  false,
			MaxIdleConns:       100,
			IdleConnTimeout:    60 * time.Second,
			DisableCompression: false,
		},
		debug)
}

// NewRestClientWithTransport creates a Rest Client with Transport
func NewRestClientWithTransport(serverAddr string, transport *http.Transport, debug bool) *RestClient {
	client := &RestClient{
		defaultTransport: transport,
	}

	client.httpClient = resty.New().
		SetTransport(client.defaultTransport).
		SetScheme(defaultHTTPSchema).
		SetAllowGetMethodPayload(false).
		SetBaseURL(fmt.Sprintf(`%s://%s/%s`, defaultHTTPSchema, serverAddr, apiURI)).
		SetTimeout(30 * time.Second).
		SetDebug(debug).
		EnableTrace()

	return client
}
