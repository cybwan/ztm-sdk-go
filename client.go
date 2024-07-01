package ztm_go_sdk

import (
	"net/http"
)

// NewClient creates a ZTM Client
func NewClient(caServerAddr, agentServerAddr string, hubServerAddr ...string) *Client {
	client := &Client{
		ca: &CaClient{
			NewRestClient(caServerAddr),
		},
		agent: &AgentClient{
			NewRestClient(agentServerAddr),
		},
		hubServerAddrs: hubServerAddr,
	}
	return client
}

// NewClientWithTransport creates a ZTM Client with Transport
func NewClientWithTransport(caServerAddr, agentServerAddr string, transport *http.Transport, hubServerAddr ...string) *Client {
	client := &Client{
		ca: &CaClient{
			NewRestClientWithTransport(caServerAddr, transport),
		},
		agent: &AgentClient{
			NewRestClientWithTransport(agentServerAddr, transport),
		},
		hubServerAddrs: hubServerAddr,
	}
	return client
}

func (c *Client) Invite(username string) (*Permit, error) {
	var err error
	permit := new(Permit)
	if permit.CA, err = c.ca.Ca(); err != nil {
		return nil, err
	}
	if permit.Agent.PrivateKey, err = c.ca.PrivateKey(username); err != nil {
		return nil, err
	}
	if permit.Agent.Certificate, err = c.ca.Certificate(username); err != nil {
		return nil, err
	}
	permit.Bootstraps = c.hubServerAddrs
	return permit, nil
}

func (c *Client) Evict(username string) (bool, error) {
	return c.ca.Delete(username)
}

func (c *Client) JoinMesh(meshname, epname string, perm *Permit) error {
	return c.agent.JoinMesh(meshname, epname, perm)
}

func (c *Client) ListMeshes() ([]*Mesh, error) {
	return c.agent.ListMeshes()
}

func (c *Client) GetMesh(meshname string) (*Mesh, error) {
	return c.agent.GetMesh(meshname)
}
