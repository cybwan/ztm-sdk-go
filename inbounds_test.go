package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListInbounds(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	outbounds, _ := client.ListInbounds(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL)
	bytes, _ := json.MarshalIndent(outbounds, "", " ")
	fmt.Println(string(bytes))
}

func TestDescribeInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	inbound, _ := client.DescribeInbound(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, TCP, "httpbin")
	bytes, _ := json.MarshalIndent(inbound, "", " ")
	fmt.Println(string(bytes))
}

func TestOpenInbound(t *testing.T) {
	client := InboundClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	err := client.OpenInbound(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, TCP, "httpbin", []Listen{
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
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	err := client.CloseInbound(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, TCP, "httpbin")
	if err != nil {
		t.Error(err)
	}
}
