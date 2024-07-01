package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	ztm "github.com/cybwan/ztm-go-sdk"
)

func main() {
	meshname := `k8s`

	homeUser := `fsm-home`
	officeUser := `fsm-office`

	homeZtmClient := ztm.NewClient(`192.168.226.61:9999`, `192.168.226.63:7777`, `192.168.226.62:8888`)
	officeZtmClient := ztm.NewClient(`192.168.226.61:9999`, `192.168.226.64:7777`, `192.168.226.62:8888`)

	homeUserPerm, _ := loadPermit(homeZtmClient, homeUser)
	officeUserPerm, _ := loadPermit(homeZtmClient, officeUser)

	if err := homeZtmClient.Join(meshname, "cluster-home", homeUserPerm); err != nil {
		fmt.Println(err.Error())
	}

	if err := officeZtmClient.Join(meshname, "cluster-office", officeUserPerm); err != nil {
		fmt.Println(err.Error())
	}

	var homeMesh *ztm.Mesh = nil
	var officeMesh *ztm.Mesh = nil

	if meshes, err := homeZtmClient.ListMeshes(); err == nil {
		fmt.Println("ListMeshes of Home")
		for idx, mesh := range meshes {
			fmt.Println(idx, mesh.MeshName, mesh.Agent.ID, mesh.Agent.EndpointName, mesh.Agent.UserName)
		}
		if len(meshes) > 0 {
			homeMesh = meshes[0]
		}
	}

	if meshes, err := officeZtmClient.ListMeshes(); err == nil {
		fmt.Println("ListMeshes of Office")
		for idx, mesh := range meshes {
			fmt.Println(idx, mesh.MeshName, mesh.Agent.ID, mesh.Agent.EndpointName, mesh.Agent.UserName)
		}
		if len(meshes) > 0 {
			officeMesh = meshes[0]
		}
	}

	var homeEndpoint *ztm.MeshEndpoint = nil
	var officeEndpoint *ztm.MeshEndpoint = nil

	if meshEndpoints, err := homeZtmClient.ListMeshEndpoints(homeMesh.MeshName); err == nil {
		fmt.Println("")
		fmt.Println("ListMeshEndpoints of home")
		for idx, meshEndpoint := range meshEndpoints {
			fmt.Println(idx, meshEndpoint.UUID, meshEndpoint.Name, meshEndpoint.IP, meshEndpoint.Port, meshEndpoint.Online, meshEndpoint.IsLocal)
			if meshEndpoint.IsLocal {
				homeEndpoint = meshEndpoint
			}
		}
	}

	if meshEndpoints, err := officeZtmClient.ListMeshEndpoints(officeMesh.MeshName); err == nil {
		fmt.Println("")
		fmt.Println("ListMeshEndpoints of office")
		for idx, meshEndpoint := range meshEndpoints {
			fmt.Println(idx, meshEndpoint.UUID, meshEndpoint.Name, meshEndpoint.IP, meshEndpoint.Port, meshEndpoint.Online, meshEndpoint.IsLocal)
			if meshEndpoint.IsLocal {
				officeEndpoint = meshEndpoint
			}
		}
	}

	if err := homeZtmClient.CreateEndpointService(homeMesh.MeshName, homeEndpoint.UUID, ztm.TCP, "nginx", "127.0.0.1", 80); err != nil {
		fmt.Println(err.Error())
	}

	if err := officeZtmClient.CreateEndpointPort(officeMesh.MeshName, officeEndpoint.UUID, ztm.TCP, "192.168.127.64", 10064, "nginx"); err != nil {
		fmt.Println(err.Error())
	}

	if err := officeZtmClient.CreateEndpointService(officeMesh.MeshName, officeEndpoint.UUID, ztm.TCP, "nginx", "127.0.0.1", 80); err != nil {
		fmt.Println(err.Error())
	}

	if err := homeZtmClient.CreateEndpointPort(homeMesh.MeshName, homeEndpoint.UUID, ztm.TCP, "192.168.127.63", 10063, "nginx"); err != nil {
		fmt.Println(err.Error())
	}
}

func loadPermit(ztmClient *ztm.Client, user string) (*ztm.Permit, error) {
	permfile := fmt.Sprintf("perms/%s.json", user)
	if _, statErr := os.Stat(permfile); statErr == nil {
		file, fileErr := os.Open(permfile)
		if fileErr != nil {
			return nil, fileErr
		}
		defer file.Close()

		bytes, readErr := io.ReadAll(file)
		if readErr != nil {
			return nil, readErr
		}
		perm := new(ztm.Permit)
		if err := json.Unmarshal(bytes, perm); err != nil {
			return nil, err
		}
		return perm, nil
	}

	ztmClient.Evict(user)
	permit, apiErr := ztmClient.Invite(user)
	if apiErr != nil {
		return nil, apiErr
	}

	bytes, _ := json.MarshalIndent(permit, "", "")
	if file, fileErr := os.OpenFile(fmt.Sprintf("perms/%s.json", user), os.O_WRONLY|os.O_CREATE, 0666); fileErr == nil {
		defer file.Close()
		io.WriteString(file, string(bytes))
	}

	return permit, nil
}
