package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}

	permit, err := LoadPermit("/tmp/root.json")
	if err != nil {
		t.Error(err)
	}

	err = client.Join(currCtx.meshName, currCtx.endpointName, permit)
	if err != nil {
		t.Error(err)
	}
}

func TestListMeshes(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	meshes, _ := client.ListMeshes()
	bytes, _ := json.Marshal(meshes)
	fmt.Println(string(bytes))
}

func TestGetMesh(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	mesh, _ := client.GetMesh(currCtx.meshName)
	bytes, _ := json.Marshal(mesh)
	fmt.Println(string(bytes))
}

func TestLeave(t *testing.T) {
	client := MeshesClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}

	err := client.Leave(currCtx.meshName)
	if err != nil {
		t.Error(err)
	}
}
