package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListEndpoints(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	endpoints, _ := client.ListEndpoints(currCtx.meshName)
	bytes, _ := json.MarshalIndent(endpoints, "", " ")
	fmt.Println(string(bytes))
}

func TestGetEndpoint(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	endpoint, _ := client.GetEndpoint(currCtx.meshName, currCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(endpoint, "", " ")
	fmt.Println(string(bytes))
}

func TestEndpointLogs(t *testing.T) {
	client := EndpointClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	logs, _ := client.GetEndpointLogs(currCtx.meshName, currCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(logs, "", " ")
	fmt.Println(string(bytes))
}
