package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

type MeshesClient struct {
	*RestClient
}

func (c *MeshesClient) Join(meshName, endpointName string, perm *Permit) error {
	perm.Agent.EndpointName = endpointName
	resp, err := c.httpClient.R().
		SetBody(perm).
		Post(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *MeshesClient) Leave(meshName string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusNoContent {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *MeshesClient) ListMeshes() ([]*Mesh, error) {
	resp, err := c.httpClient.R().
		SetResult([]*Mesh{}).
		Get(fmt.Sprintf("%s", "meshes"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}

	meshes := resp.Result().(*[]*Mesh)

	return *meshes, nil
}

func (c *MeshesClient) GetMesh(meshName string) (*Mesh, error) {
	resp, err := c.httpClient.R().
		SetResult(&Mesh{}).
		Get(fmt.Sprintf("%s/%s", "meshes", meshName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}

	mesh := resp.Result().(*Mesh)

	return mesh, nil
}

func (c *MeshesClient) LocalMesh() string {
	if meshes, err := c.ListMeshes(); err == nil {
		if len(meshes) > 0 {
			return meshes[0].MeshName
		}
	}
	return ""
}
