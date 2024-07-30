package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

type OutboundClient struct {
	*RestClient
}

func (c *OutboundClient) ListOutbounds(meshName, endpointId, provider, appName string) ([]*Outbound, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Outbound{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "outbound"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	outbounds := resp.Result().(*[]*Outbound)

	return *outbounds, nil
}

func (c *OutboundClient) DescribeOutbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string) (*Outbound, error) {
	resp, err := c.httpClient.R().
		SetResult(Outbound{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "outbound", protocol, name))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	outbound := resp.Result().(*Outbound)

	return outbound, nil
}
func (c *OutboundClient) OpenOutbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string, targets []Target) error {
	resp, err := c.httpClient.R().
		SetBody(Outbound{
			Targets: targets,
		}).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "outbound", protocol, name))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *OutboundClient) CloseOutbound(meshName, endpointId, provider, appName string, protocol MeshProtocol, name string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "apps", provider, appName, "api", "endpoints", endpointId, "outbound", protocol, name))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusNoContent && resp.StatusCode() != http.StatusNotFound {
		return errors.New(resp.Status())
	}

	return nil
}
