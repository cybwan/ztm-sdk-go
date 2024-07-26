package ztm_go_sdk

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	apiURI            = "api"
	defaultHTTPSchema = "http"
)

type RestClient struct {
	defaultTransport *http.Transport
	httpClient       *resty.Client
}

type HubClient struct {
	*RestClient
}

type AgentClient struct {
	*RestClient
	*MeshesClient
	*EndpointClient
	*AppClient
	*OutboundClient
	*InboundClient
	*FileClient
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

type Endpoint struct {
	Local     bool     `json:"isLocal,omitempty"`
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

type App struct {
	Provider   string `json:"provider"`
	Name       string `json:"name"`
	Tag        string `json:"tag,omitempty"`
	Builtin    bool   `json:"isBuiltin"`
	Downloaded bool   `json:"isDownloaded"`
	Published  bool   `json:"isPublished"`
	Running    bool   `json:"isRunning"`
}

type File struct {
	Name    string   `json:"name"`
	Size    uint32   `json:"size"`
	Time    uint64   `json:"time"`
	Hash    string   `json:"hash"`
	Sources []string `json:"sources,omitempty"`
}

type Target struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

type Outbound struct {
	Name      string       `json:"name,omitempty"`
	Protocol  MeshProtocol `json:"protocol,omitempty"`
	Targets   []Target     `json:"targets,omitempty"`
	Entrances []string     `json:"entrances,omitempty"`
}

type Listen struct {
	IP   string `json:"ip"`
	Port uint16 `json:"port"`
	Open bool   `json:"open"`
}

type Inbound struct {
	Name     string       `json:"name,omitempty"`
	Protocol MeshProtocol `json:"protocol,omitempty"`
	Listens  []Listen     `json:"listens,omitempty"`
	Exits    []string     `json:"exits,omitempty"`
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
