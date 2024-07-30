package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListApps(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	apps, _ := client.ListApps(currCtx.meshName, currCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(apps, "", " ")
	fmt.Println(string(bytes))
}

func TestGetApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	app, _ := client.GetApp(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", "")
	bytes, _ := json.MarshalIndent(app, "", " ")
	fmt.Println(string(bytes))
}

func TestStartApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	_, err := client.StartApp(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", "")
	if err != nil {
		t.Error(err)
	}
}

func TestStopApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	_, err := client.StopApp(currCtx.meshName, currCtx.LocalEndpointId(), "ztm", "tunnel", "")
	if err != nil {
		t.Error(err)
	}
}
