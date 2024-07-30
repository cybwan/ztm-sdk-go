package ztm_go_sdk

import (
	"net/http"
)

// NewAgentClient creates an Agent Client
func NewAgentClient(agentServerAddr string, debug bool) *AgentClient {
	restClient := NewRestClient(agentServerAddr, debug)
	return &AgentClient{
		MeshesClient: &MeshesClient{
			RestClient: restClient,
		},
		EndpointClient: &EndpointClient{
			RestClient: restClient,
		},
		AppClient: &AppClient{
			RestClient: restClient,
		},
		OutboundClient: &OutboundClient{
			RestClient: restClient,
		},
		InboundClient: &InboundClient{
			RestClient: restClient,
		},
		FileClient: &FileClient{
			RestClient: restClient,
		},
	}
}

// NewAgentClientWithTransport creates an Agent Client with Transport
func NewAgentClientWithTransport(agentServerAddr string, transport *http.Transport, debug bool) *AgentClient {
	restClient := NewRestClientWithTransport(agentServerAddr, transport, debug)
	return &AgentClient{
		MeshesClient: &MeshesClient{
			RestClient: restClient,
		},
		EndpointClient: &EndpointClient{
			RestClient: restClient,
		},
		AppClient: &AppClient{
			RestClient: restClient,
		},
		OutboundClient: &OutboundClient{
			RestClient: restClient,
		},
		InboundClient: &InboundClient{
			RestClient: restClient,
		},
		FileClient: &FileClient{
			RestClient: restClient,
		},
	}
}
