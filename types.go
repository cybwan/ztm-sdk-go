package ztm_go_sdk

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	apiURI            = "api"
	defaultHTTPSchema = "http"
)

type Client struct {
	ca             *CaClient
	agent          *AgentClient
	hubServerAddrs []string
}

type RestClient struct {
	defaultTransport *http.Transport
	httpClient       *resty.Client
}

type CaClient struct {
	*RestClient
}

type HubClient struct {
	*RestClient
}

type AgentClient struct {
	*RestClient
}

type Permit struct {
	CA    string `json:"ca"`
	Agent struct {
		EndpointName string `json:"name,omitempty"`
		Certificate  string `json:"certificate"`
		PrivateKey   string `json:"privateKey"`
	} `json:"agent"`
	Bootstraps []string `json:"bootstraps"`
}

type Mesh struct {
	MeshName string `json:"name,omitempty"`
	CA       string `json:"ca"`
	Agent    struct {
		ID           string `json:"id,omitempty"`
		EndpointName string `json:"name,omitempty"`
		UserName     string `json:"username,omitempty"`
		Certificate  string `json:"certificate"`
		PrivateKey   string `json:"privateKey"`
	} `json:"agent"`
	Bootstraps []string `json:"bootstraps"`
	Connected  bool     `json:"connected,omitempty"`
}
