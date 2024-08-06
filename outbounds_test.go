package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListOutbounds(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	outbounds, _ := client.ListOutbounds(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel)
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}

func TestDescribeOutbound(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	outbound, _ := client.DescribeOutbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin")
	bytes, _ := json.MarshalIndent(outbound, "", " ")
	fmt.Println(string(bytes))
}

func TestOpenOutbound(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	err := client.OpenOutbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin", []Target{
		{
			Host: "44.207.203.25",
			Port: 80,
		},
		{
			Host: "54.87.89.151",
			Port: 80,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCloseOutbound(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	err := client.CloseOutbound(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, TCP, "httpbin")
	if err != nil {
		t.Error(err)
	}
}
