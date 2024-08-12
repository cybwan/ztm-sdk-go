package tunnel

import (
	ztm "github.com/cybwan/ztm-sdk-go"
)

type Client struct {
	OutboundClient *ztm.OutboundClient
	InboundClient  *ztm.InboundClient
}

func (c *Client) ListOutbounds(meshName, endpointId, provider string) ([]*ztm.Outbound, error) {
	return c.OutboundClient.ListOutbounds(meshName, endpointId, provider, APP)
}

func (c *Client) DescribeOutbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string) (*ztm.Outbound, error) {
	return c.OutboundClient.DescribeOutbound(meshName, endpointId, provider, APP, protocol, name)
}

func (c *Client) OpenOutbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string, targets []ztm.Target) error {
	return c.OutboundClient.OpenOutbound(meshName, endpointId, provider, APP, protocol, name, targets)
}

func (c *Client) CloseOutbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string) error {
	return c.OutboundClient.CloseOutbound(meshName, endpointId, provider, APP, protocol, name)
}

func (c *Client) ListInbounds(meshName, endpointId, provider string) ([]*ztm.Inbound, error) {
	return c.InboundClient.ListInbounds(meshName, endpointId, provider, APP)
}

func (c *Client) DescribeInbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string) (*ztm.Inbound, error) {
	return c.InboundClient.DescribeInbound(meshName, endpointId, provider, APP, protocol, name)
}
func (c *Client) OpenInbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string, listens []ztm.Listen) error {
	return c.InboundClient.OpenInbound(meshName, endpointId, provider, APP, protocol, name, listens)
}

func (c *Client) CloseInbound(meshName, endpointId, provider string, protocol ztm.MeshProtocol, name string) error {
	return c.InboundClient.CloseInbound(meshName, endpointId, provider, APP, protocol, name)
}
