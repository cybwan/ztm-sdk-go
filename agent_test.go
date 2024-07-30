package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAgent(t *testing.T) {
	client := NewAgentClient(currCtx.agentAddr, false)
	meshes, _ := client.ListMeshes()
	bytes, _ := json.Marshal(meshes)
	fmt.Println(string(bytes))
}
