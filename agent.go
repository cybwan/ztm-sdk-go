package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *AgentClient) ListMeshes() ([]*Mesh, error) {
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

func (c *AgentClient) GetMesh(meshname string) (*Mesh, error) {
	resp, err := c.httpClient.R().
		SetResult(&Mesh{}).
		Get(fmt.Sprintf("%s/%s", "meshes", meshname))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.Status())
	}

	mesh := resp.Result().(*Mesh)

	return mesh, nil
}

func (c *AgentClient) JoinMesh(meshname, epname string, perm *Permit) error {
	perm.Agent.EndpointName = epname
	resp, err := c.httpClient.R().
		SetBody(perm).
		Post(fmt.Sprintf("%s/%s", "meshes", meshname))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	return nil
}
