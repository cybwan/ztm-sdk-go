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
	ca    *CaClient
	agent *AgentClient
}

type RestClient struct {
	defaultTransport *http.Transport
	httpClient       *resty.Client
}

type CaClient struct {
	*RestClient
	hubServerAddrs []string
}

type HubClient struct {
	*RestClient
}

type AgentClient struct {
	*RestClient
}

type Permit struct {
	CA    string `json:"ca,omitempty"`
	Agent struct {
		EndpointName string `json:"name,omitempty"`
		Certificate  string `json:"certificate,omitempty"`
		PrivateKey   string `json:"privateKey,omitempty"`
	} `json:"agent,omitempty"`
	Bootstraps []string `json:"bootstraps,omitempty"`
}

type Mesh struct {
	MeshName string `json:"name,omitempty"`
	CA       string `json:"ca,omitempty"`
	Agent    struct {
		ID           string `json:"id,omitempty"`
		EndpointName string `json:"name,omitempty"`
		UserName     string `json:"username,omitempty"`
		Certificate  string `json:"certificate,omitempty"`
		PrivateKey   string `json:"privateKey,omitempty"`
	} `json:"agent,omitempty"`
	Bootstraps []string `json:"bootstraps,omitempty"`
	Connected  bool     `json:"connected,omitempty"`
}

type MeshEndpoint struct {
	IsLocal   bool     `json:"isLocal,omitempty"`
	UUID      string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	UserName  string   `json:"username,omitempty"`
	IP        string   `json:"ip,omitempty"`
	Port      uint16   `json:"port,omitempty"`
	Hubs      []string `json:"hubs,omitempty"`
	Users     []string `json:"users,omitempty"`
	Heartbeat uint64   `json:"heartbeat,omitempty"`
	Online    bool     `json:"online,omitempty"`
}

type MeshService struct {
	Name         string         `json:"name,omitempty"`
	Protocol     MeshProtocol   `json:"protocol,omitempty"`
	Endpoints    []MeshEndpoint `json:"endpoints,omitempty"`
	Users        []string       `json:"users,omitempty"`
	IsDiscovered bool           `json:"isDiscovered,omitempty"`
	IsLocal      bool           `json:"isLocal,omitempty"`
	Host         string         `json:"host,omitempty"`
	Port         uint16         `json:"port,omitempty"`
}

type MeshServerAddr struct {
	Host string `json:"host,omitempty"`
	Port uint16 `json:"port,omitempty"`
}

type MeshListen struct {
	IP   string `json:"ip,omitempty"`
	Port uint16 `json:"port,omitempty"`
}

type MeshTarget struct {
	Service string `json:"service,omitempty"`
}

type MeshPort struct {
	Protocol MeshProtocol `json:"protocol,omitempty"`
	Listen   MeshListen   `json:"listen,omitempty"`
	Target   MeshTarget   `json:"target,omitempty"`
	Open     bool         `json:"open,omitempty"`
}

type MeshPortTarget struct {
	Target MeshTarget `json:"target,omitempty"`
}

type Log struct {
	Time    string `json:"time,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type MeshProtocol string

const (
	TCP MeshProtocol = `tcp`
	UDP MeshProtocol = `udp`
)
