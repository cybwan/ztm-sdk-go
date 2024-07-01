package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	ztm "github.com/cybwan/ztm-go-sdk"
)

func main() {
	homeZtmClient := ztm.NewClient(`192.168.226.61:9999`, `192.168.226.63:7777`, `192.168.226.62:8888`)

	//homePerm, _ := loadPermit(homeZtmClient, `home`)
	//officePerm, _ := loadPermit(homeZtmClient, `office`)

	//err := homeZtmClient.JoinMesh("myhub", "home", homePerm)
	//fmt.Println(err)

	//meshes, err := homeZtmClient.ListMeshes()
	//fmt.Println(err)
	//fmt.Println(meshes[0].Agent.ID)

	mesh, err := homeZtmClient.GetMesh("myhub")
	fmt.Println(err)
	fmt.Println(mesh.Agent.ID)
	fmt.Println(mesh.Agent.UserName)
	fmt.Println(mesh.Agent.EndpointName)

	//fmt.Println(officePerm)
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
