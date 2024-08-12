package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListInbounds(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	outbounds, _ := client.ListInbounds(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel)
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}

func TestDescribeInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	inbound, _ := client.DescribeInbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin")
	bytes, _ := json.MarshalIndent(inbound, "", " ")
	fmt.Println(string(bytes))
}

func TestOpenInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	err := client.OpenInbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin", []Listen{
		{
			IP:   CurrCtx.hostIP,
			Port: 10081,
		},
		{
			IP:   CurrCtx.hostIP,
			Port: 10082,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCloseInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	err := client.CloseInbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin")
	if err != nil {
		t.Error(err)
	}
}
