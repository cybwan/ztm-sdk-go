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

func (c *Client) Invite(userName string) (*Permit, error) {
	var err error
	permit := new(Permit)
	if permit.CA, err = c.ca.Ca(); err != nil {
		return nil, err
	}
	if permit.Agent.PrivateKey, err = c.ca.PrivateKey(userName); err != nil {
		return nil, err
	}
	if permit.Agent.Certificate, err = c.ca.Certificate(userName); err != nil {
		return nil, err
	}
	permit.Bootstraps = c.hubServerAddrs
	return permit, nil
}

func (c *Client) Evict(userName string) (bool, error) {
	return c.ca.Delete(userName)
}

func (c *Client) Join(meshName, endpointName string, perm *Permit) error {
	return c.agent.join(meshName, endpointName, perm)
}

func (c *Client) ListMeshes() ([]*Mesh, error) {
	return c.agent.listMeshes()
}

func (c *Client) GetMesh(meshName string) (*Mesh, error) {
	return c.agent.getMesh(meshName)
}

func (c *Client) ListMeshEndpoints(meshName string) ([]*MeshEndpoint, error) {
	return c.agent.listMeshEndpoints(meshName)
}

func (c *Client) GetMeshEndpoint(meshName, endpointUuid string) (*MeshEndpoint, error) {
	return c.agent.getMeshEndpoint(meshName, endpointUuid)
}

func (c *Client) GetMeshEndpointLogs(meshName, endpointUuid string) ([]*Log, error) {
	return c.agent.getMeshEndpointLogs(meshName, endpointUuid)
}

func (c *Client) ListMeshServices(meshName string) ([]*MeshService, error) {
	return c.agent.listMeshServices(meshName)
}

func (c *Client) GetMeshService(meshName string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
	return c.agent.getMeshService(meshName, protocol, serviceName)
}

func (c *Client) ListEndpointServices(meshName, endpointUuid string) ([]*MeshService, error) {
	return c.agent.listEndpointServices(meshName, endpointUuid)
}

func (c *Client) GetEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
	return c.agent.getEndpointService(meshName, endpointUuid, protocol, serviceName)
}

func (c *Client) CreateEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName, host string, port uint16) error {
	return c.agent.createEndpointService(meshName, endpointUuid, protocol, serviceName, MeshServerAddr{Host: host, Port: port})
}

func (c *Client) DeleteEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) error {
	return c.agent.deleteEndpointService(meshName, endpointUuid, protocol, serviceName)
}

func (c *Client) ListEndpointPorts(meshName, endpointUuid string) ([]*MeshPort, error) {
	return c.agent.listEndpointPorts(meshName, endpointUuid)
}

func (c *Client) GetEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) (*MeshPort, error) {
	return c.agent.getEndpointPort(meshName, endpointUuid, protocol, ip, port)
}

func (c *Client) CreateEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16, targetService string) error {
	return c.agent.createEndpointPort(meshName, endpointUuid, protocol, ip, port, MeshPortTarget{Target: MeshTarget{Service: targetService}})
}

func (c *Client) DeleteEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) error {
	return c.agent.deleteEndpointPort(meshName, endpointUuid, protocol, ip, port)
}
