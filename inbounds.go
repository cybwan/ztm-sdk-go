package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

type InboundClient struct {
	*RestClient
}

func (c *InboundClient) ListInbounds(meshName, endpointId, provider, appName string) ([]*Inbound, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Inbound{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "inbound"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	inbounds := resp.Result().(*[]*Inbound)

	return *inbounds, nil
}

func (c *InboundClient) DescribeInbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string) (*Inbound, error) {
	resp, err := c.httpClient.R().
		SetResult(Inbound{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "inbound", protocol, name))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	inbound := resp.Result().(*Inbound)

	return inbound, nil
}
func (c *InboundClient) OpenInbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string, listens []Listen) error {
	resp, err := c.httpClient.R().
		SetBody(Inbound{
			Listens: listens,
		}).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "inbound", protocol, name))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *InboundClient) CloseInbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "inbound", protocol, name))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusNoContent && resp.StatusCode() != http.StatusNotFound {
		return errors.New(resp.Status())
	}

	return nil
}
