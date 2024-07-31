package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListApps(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	apps, _ := client.ListApps(currCtx.meshName, currCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(apps, "", " ")
	fmt.Println(string(bytes))
}

func TestGetApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	app, _ := client.GetApp(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, "")
	bytes, _ := json.MarshalIndent(app, "", " ")
	fmt.Println(string(bytes))
}

func TestStartApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	_, err := client.StartApp(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, "")
	if err != nil {
		t.Error(err)
	}
}

func TestStopApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(currCtx.agentAddr, false),
	}
	_, err := client.StopApp(currCtx.meshName, currCtx.LocalEndpointId(), ZTM, APP_TUNNEL, "")
	if err != nil {
		t.Error(err)
	}
}
