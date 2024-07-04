package ztm_go_sdk

import (
	"net/http"
)

// NewClient creates a ZTM Client
func NewClient(caServerAddr, agentServerAddr string, hubServerAddr ...string) *Client {
	client := &Client{
		ca:    NewCaClient(caServerAddr, hubServerAddr...),
		agent: NewAgentClient(agentServerAddr),
	}
	return client
}

// NewClientWithTransport creates a ZTM Client with Transport
func NewClientWithTransport(caServerAddr, agentServerAddr string, transport *http.Transport, hubServerAddr ...string) *Client {
	client := &Client{
		ca:    NewCaClientWithTransport(caServerAddr, transport, hubServerAddr...),
		agent: NewAgentClientWithTransport(agentServerAddr, transport),
	}
	return client
}

func (c *Client) Ca() *CaClient {
	return c.ca
}

func (c *Client) Agent() *AgentClient {
	return c.agent
}
