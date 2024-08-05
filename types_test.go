package ztm_go_sdk

import (
	"os"
	"strings"
)

var (
	homeCtx = &ZtmContext{
		user:         "root",
		agentAddr:    "127.0.0.1:7771",
		meshName:     "k8s-mesh",
		endpointName: "home",
	}

	officeCtx = &ZtmContext{
		user:         "root",
		agentAddr:    "127.0.0.1:7772",
		meshName:     "k8s-mesh",
		endpointName: "office",
	}

	currCtx *ZtmContext

	APP_TUNNEL = "tunnel"
)

type ZtmContext struct {
	user         string
	agentAddr    string
	meshName     string
	endpointName string
	hostIP       string
}

func (c *ZtmContext) LocalEndpointId() string {
	client := EndpointClient{
		RestClient: NewRestClient(c.agentAddr, false),
	}

	if localEndpoint, err := client.LocalEndpoint(c.meshName); err == nil {
		return localEndpoint.UUID
	}

	return c.endpointName
}

func init() {
	agent := os.Getenv("CTR_AGENT")
	if strings.EqualFold(agent, "office") {
		currCtx = officeCtx
	} else {
		currCtx = homeCtx
	}

	hostIP := os.Getenv("MY_HOST_IP")
	if len(hostIP) == 0 {
		hostIP = "127.0.0.1"
	}
	currCtx.hostIP = hostIP
}
