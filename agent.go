package ztm_go_sdk

import (
	"net/http"
)

// NewAgentClient creates an Agent Client
func NewAgentClient(agentServerAddr string) *AgentClient {
	return &AgentClient{
		RestClient: NewRestClient(agentServerAddr),
	}
}

// NewAgentClientWithTransport creates an Agent Client with Transport
func NewAgentClientWithTransport(agentServerAddr string, transport *http.Transport) *AgentClient {
	return &AgentClient{
		RestClient: NewRestClientWithTransport(agentServerAddr, transport),
	}
}
