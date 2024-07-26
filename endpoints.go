package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

type EndpointClient struct {
	*RestClient
}

func (c *EndpointClient) ListEndpoints(meshName string) ([]*Endpoint, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Endpoint{}).
		Get(fmt.Sprintf("%s/%s/%s", "meshes", meshName, "endpoints"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshEndpoints := resp.Result().(*[]*Endpoint)

	return *meshEndpoints, nil
}

func (c *EndpointClient) GetEndpoint(meshName, endpointId string) (*Endpoint, error) {
	resp, err := c.httpClient.R().
		SetResult(Endpoint{}).
		Get(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointId))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	meshEndpoint := resp.Result().(*Endpoint)

	return meshEndpoint, nil
}

func (c *EndpointClient) GetEndpointLogs(meshName, endpointUuid string) ([]*Log, error) {
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

func (c *EndpointClient) LocalEndpoint(meshName string) (*Endpoint, error) {
	endpoints, err := c.ListEndpoints(meshName)
	if err == nil {
		for _, endpoint := range endpoints {
			if endpoint.Local {
				return endpoint, nil
			}
		}
	}
	return nil, err
}
