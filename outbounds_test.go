package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListOutbounds(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	outbounds, _ := client.ListOutbounds(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel")
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}

func TestDescribeOutbound(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	outbound, _ := client.DescribeOutbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin")
	bytes, _ := json.MarshalIndent(outbound, "", " ")
	fmt.Println(string(bytes))
}

func TestOpenOutbound(t *testing.T) {
	client := OutboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	err := client.OpenOutbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin", []Target{
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
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	err := client.CloseOutbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin")
	if err != nil {
		t.Error(err)
	}
}
