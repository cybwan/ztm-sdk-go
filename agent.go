package ztm_go_sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (c *AgentClient) ListMeshes() ([]*Mesh, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Mesh{}).
		Get(fmt.Sprintf("%s", "meshes"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}

	meshes := resp.Result().(*[]*Mesh)

	return *meshes, nil
}

func (c *AgentClient) GetMesh(meshName string) (*Mesh, error) {
	resp, err := c.httpClient.R().
		SetResult(&Mesh{}).
		Get(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}

	mesh := resp.Result().(*Mesh)

	return mesh, nil
}

func (c *AgentClient) Join(meshName, endpointName string, perm *Permit) error {
	perm.Agent.EndpointName = endpointName
	resp, err := c.httpClient.R().
		SetBody(perm).
		Post(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) Leave(meshName string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusNoContent {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) ListMeshEndpoints(meshName string) ([]*MeshEndpoint, error) {
	resp, err := c.httpClient.R().
		SetResult([]*MeshEndpoint{}).
		Get(fmt.Sprintf("%s/%s/%s", "meshes", meshName, "endpoints"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshEndpoints := resp.Result().(*[]*MeshEndpoint)

	return *meshEndpoints, nil
}

func (c *AgentClient) GetMeshEndpoint(meshName, endpointUuid string) (*MeshEndpoint, error) {
	resp, err := c.httpClient.R().
		SetResult(MeshEndpoint{}).
		Get(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshEndpoint := resp.Result().(*MeshEndpoint)

	return meshEndpoint, nil
}

func (c *AgentClient) GetMeshEndpointLogs(meshName, endpointUuid string) ([]*Log, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Log{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "log"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	logs := resp.Result().(*[]*Log)

	return *logs, nil
}

func (c *AgentClient) ListMeshServices(meshName string) ([]*MeshService, error) {
	resp, err := c.httpClient.R().
		SetResult([]*MeshService{}).
		Get(fmt.Sprintf("%s/%s/%s", "meshes", meshName, "services"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshServices := resp.Result().(*[]*MeshService)

	return *meshServices, nil
}

func (c *AgentClient) GetMeshService(meshName string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
	resp, err := c.httpClient.R().
		SetResult(MeshService{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s", "meshes", meshName, "services", protocol, serviceName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshService := resp.Result().(*MeshService)

	return meshService, nil
}

func (c *AgentClient) ListEndpointServices(meshName, endpointUuid string) ([]*MeshService, error) {
	resp, err := c.httpClient.R().
		SetResult([]*MeshService{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "services"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshServices := resp.Result().(*[]*MeshService)

	return *meshServices, nil
}

func (c *AgentClient) GetEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
	resp, err := c.httpClient.R().
		SetResult(MeshService{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "services", protocol, serviceName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshService := resp.Result().(*MeshService)

	return meshService, nil
}

func (c *AgentClient) CreateEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName, host string, port uint16) error {
	resp, err := c.httpClient.R().
		SetBody(MeshServerAddr{Host: host, Port: port}).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "services", protocol, serviceName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) DeleteEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "services", protocol, serviceName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) ListEndpointPorts(meshName, endpointUuid string) ([]*MeshPort, error) {
	resp, err := c.httpClient.R().
		SetResult([]*MeshPort{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "ports"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshPorts := resp.Result().(*[]*MeshPort)

	return *meshPorts, nil
}

func (c *AgentClient) GetEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) (*MeshPort, error) {
	resp, err := c.httpClient.R().
		SetResult(MeshPort{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%d", "meshes", meshName, "endpoints", endpointUuid, "ports", ip, protocol, port))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshPort := resp.Result().(*MeshPort)

	return meshPort, nil
}

func (c *AgentClient) CreateEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16, targetService string) error {
	resp, err := c.httpClient.R().
		SetBody(MeshPortTarget{Target: MeshTarget{Service: targetService}}).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%d", "meshes", meshName, "endpoints", endpointUuid, "ports", ip, protocol, port))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) DeleteEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%d", "meshes", meshName, "endpoints", endpointUuid, "ports", ip, protocol, port))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) LoadPermit(permfile string) (*Permit, error) {
	if _, statErr := os.Stat(permfile); statErr != nil {
		return nil, statErr
	}
	file, fileErr := os.Open(permfile)
	if fileErr != nil {
		return nil, fileErr
	}
	defer file.Close()

	bytes, readErr := io.ReadAll(file)
	if readErr != nil {
		return nil, readErr
	}
	perm := new(Permit)
	if err := json.Unmarshal(bytes, perm); err != nil {
		return nil, err
	}
	return perm, nil
}
