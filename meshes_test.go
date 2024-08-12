package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}

	permit, err := LoadPermit("/tmp/.ztm.root.json")
	if err != nil {
		t.Error(err)
	}

	err = client.Join(CurrCtx.meshName, CurrCtx.endpointName, permit)
	if err != nil {
		t.Error(err)
	}
}

func TestListMeshes(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	meshes, _ := client.ListMeshes()
	bytes, _ := json.Marshal(meshes)
	fmt.Println(string(bytes))
}

func TestGetMesh(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	mesh, _ := client.GetMesh(CurrCtx.meshName)
	bytes, _ := json.Marshal(mesh)
	fmt.Println(string(bytes))
}

func TestLeave(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}

	err := client.Leave(CurrCtx.meshName)
	if err != nil {
		t.Error(err)
	}
}
