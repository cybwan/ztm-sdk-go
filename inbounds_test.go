package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListInbounds(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	outbounds, _ := client.ListInbounds(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel")
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}

func TestDescribeInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	inbound, _ := client.DescribeInbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin")
	bytes, _ := json.MarshalIndent(inbound, "", " ")
	fmt.Println(string(bytes))
}

func TestOpenInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	err := client.OpenInbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin", []Listen{
		{
			IP:   currCtx.hostIP,
			Port: 10081,
		},
		{
			IP:   currCtx.hostIP,
			Port: 10082,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCloseInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	err := client.CloseInbound(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", TCP, "httpbin")
	if err != nil {
		t.Error(err)
	}
}
