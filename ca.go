package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

// NewCaClient creates a Ca Client
func NewCaClient(caServerAddr string, hubServerAddr ...string) *CaClient {
	return &CaClient{
		RestClient:     NewRestClient(caServerAddr),
		hubServerAddrs: hubServerAddr,
	}
}

// NewCaClientWithTransport creates a Ca Client with Transport
func NewCaClientWithTransport(caServerAddr string, transport *http.Transport, hubServerAddr ...string) *CaClient {
	return &CaClient{
		RestClient:     NewRestClientWithTransport(caServerAddr, transport),
		hubServerAddrs: hubServerAddr,
	}
}

func (c *CaClient) GetCa() (string, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%s", "certificates", "ca"))

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.Status())
	}

	return string(resp.Body()), nil
}

func (c *CaClient) PrivateKey(username string) (string, error) {
	resp, err := c.httpClient.R().Post(fmt.Sprintf("%s/%s", "certificates", username))

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.Status())
	}

	return string(resp.Body()), nil
}

func (c *CaClient) Certificate(username string) (string, error) {
	resp, err := c.httpClient.R().Get(fmt.Sprintf("%s/%s", "certificates", username))

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.Status())
	}

	return string(resp.Body()), nil
}

func (c *CaClient) Delete(username string) (bool, error) {
	resp, err := c.httpClient.R().Delete(fmt.Sprintf("%s/%s", "certificates", username))

	if err != nil {
		return false, err
	}

	if resp.StatusCode() == http.StatusOK || resp.StatusCode() == http.StatusNoContent {
		return true, nil
	}

	return false, errors.New(resp.Status())
}

func (c *CaClient) Invite(userName string) (*Permit, error) {
	var err error
	permit := new(Permit)
	if permit.CA, err = c.GetCa(); err != nil {
		return nil, err
	}
	if permit.Agent.PrivateKey, err = c.PrivateKey(userName); err != nil {
		return nil, err
	}
	if permit.Agent.Certificate, err = c.Certificate(userName); err != nil {
		return nil, err
	}
	permit.Bootstraps = c.hubServerAddrs
	return permit, nil
}

func (c *CaClient) Evict(userName string) (bool, error) {
	return c.Delete(userName)
}
