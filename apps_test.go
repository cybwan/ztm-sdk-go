package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListApps(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	apps, _ := client.ListApps(CurrCtx.meshName, CurrCtx.LocalEndpointId())
	bytes, _ := json.MarshalIndent(apps, "", " ")
	fmt.Println(string(bytes))
}

func TestGetApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	app, _ := client.GetApp(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, "")
	bytes, _ := json.MarshalIndent(app, "", " ")
	fmt.Println(string(bytes))
}

func TestStartApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	_, err := client.StartApp(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, "")
	if err != nil {
		t.Error(err)
	}
}

func TestStopApp(t *testing.T) {
	client := AppClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	_, err := client.StopApp(CurrCtx.meshName, CurrCtx.LocalEndpointId(), ZTM, AppTunnel, "")
	if err != nil {
		t.Error(err)
	}
}
