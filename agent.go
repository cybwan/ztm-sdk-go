package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *AgentClient) listMeshes() ([]*Mesh, error) {
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

func (c *AgentClient) getMesh(meshName string) (*Mesh, error) {
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

func (c *AgentClient) join(meshName, endpointName string, perm *Permit) error {
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

func (c *AgentClient) leave(meshName string) error {
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

func (c *AgentClient) listMeshEndpoints(meshName string) ([]*MeshEndpoint, error) {
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

func (c *AgentClient) getMeshEndpoint(meshName, endpointUuid string) (*MeshEndpoint, error) {
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

func (c *AgentClient) getMeshEndpointLogs(meshName, endpointUuid string) ([]*Log, error) {
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

func (c *AgentClient) listMeshServices(meshName string) ([]*MeshService, error) {
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

func (c *AgentClient) getMeshService(meshName string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
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

func (c *AgentClient) listEndpointServices(meshName, endpointUuid string) ([]*MeshService, error) {
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

func (c *AgentClient) getEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) (*MeshService, error) {
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

func (c *AgentClient) createEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string, serverAddr MeshServerAddr) error {
	resp, err := c.httpClient.R().
		SetBody(serverAddr).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointUuid, "services", protocol, serviceName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) deleteEndpointService(meshName, endpointUuid string, protocol MeshProtocol, serviceName string) error {
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

func (c *AgentClient) listEndpointPorts(meshName, endpointUuid string) ([]*MeshPort, error) {
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

func (c *AgentClient) getEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) (*MeshPort, error) {
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

func (c *AgentClient) createEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16, target MeshPortTarget) error {
	resp, err := c.httpClient.R().
		SetBody(target).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%d", "meshes", meshName, "endpoints", endpointUuid, "ports", ip, protocol, port))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *AgentClient) deleteEndpointPort(meshName, endpointUuid string, protocol MeshProtocol, ip string, port uint16) error {
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
