package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

type AppClient struct {
	*RestClient
}

func (c *AppClient) ListApps(meshName, endpointId string) ([]*App, error) {
	resp, err := c.httpClient.R().
		SetResult([]*App{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointId, "apps"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	apps := resp.Result().(*[]*App)

	return *apps, nil
}

func (c *AppClient) GetApp(meshName, endpointId, provider, name, tag string) (*App, error) {
	if len(tag) > 0 {
		name = fmt.Sprintf("%s@%s", name, tag)
	}
	resp, err := c.httpClient.R().
		SetResult(App{}).
		Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointId, "apps", provider, name))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	app := resp.Result().(*App)

	return app, nil
}

func (c *AppClient) StartApp(meshName, endpointId, provider, name, tag string) (*App, error) {
	app, err := c.GetApp(meshName, endpointId, provider, name, tag)
	if err != nil {
		return nil, err
	}

	if app.Running {
		return app, nil
	}

	app.Running = true

	if len(tag) > 0 {
		name = fmt.Sprintf("%s@%s", name, tag)
	}

	resp, err := c.httpClient.R().
		SetResult(App{}).
		SetBody(app).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointId, "apps", provider, name))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(resp.Status())
	}

	app = resp.Result().(*App)

	return app, nil
}

func (c *AppClient) StopApp(meshName, endpointId, provider, name, tag string) (*App, error) {
	app, err := c.GetApp(meshName, endpointId, provider, name, tag)
	if err != nil {
		return nil, err
	}

	if !app.Running {
		return app, nil
	}

	app.Running = false

	if len(tag) > 0 {
		name = fmt.Sprintf("%s@%s", name, tag)
	}

	resp, err := c.httpClient.R().
		SetResult(App{}).
		SetBody(app).
		Post(fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", "meshes", meshName, "endpoints", endpointId, "apps", provider, name))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(resp.Status())
	}

	app = resp.Result().(*App)

	return app, nil
}
