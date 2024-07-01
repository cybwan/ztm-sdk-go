package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *CaClient) Ca() (string, error) {
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
