package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListEndpoints(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	endpoints, _ := client.ListEndpoints(CurrCtx.meshName)
	bytes, _ := json.MarshalIndent(endpoints, "", " ")
	fmt.Println(string(bytes))
}

func TestGetEndpoint(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	endpoint, _ := client.GetEndpoint(CurrCtx.meshName, CurrCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(endpoint, "", " ")
	fmt.Println(string(bytes))
}

func TestEndpointLogs(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	logs, _ := client.GetEndpointLogs(CurrCtx.meshName, CurrCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(logs, "", " ")
	fmt.Println(string(bytes))
}
