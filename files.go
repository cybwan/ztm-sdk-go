package ztm_go_sdk

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type FileClient struct {
	*RestClient
}

func (c *FileClient) ListFiles(meshName string) ([]*File, error) {
	resp, err := c.httpClient.R().
		SetResult(map[string]*File{}).
		Get(fmt.Sprintf("%s/%s/%s", "meshes", meshName, "files"))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	maps := resp.Result().(*map[string]*File)

	var files []*File

	if maps != nil && len(*maps) > 0 {
		for name, file := range *maps {
			file.Name = name
			files = append(files, file)
		}
	}

	return files, nil
}

func (c *FileClient) DescribeFile(meshName, fileName string) (*File, error) {
	resp, err := c.httpClient.R().
		SetResult(File{}).
		Get(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "files", strings.TrimPrefix(fileName, "/")))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	file := resp.Result().(*File)
	file.Name = fileName

	return file, nil
}

func (c *FileClient) DownloadFile(meshName, fileName string) (string, error) {
	resp, err := c.httpClient.R().
		Get(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "file-data", strings.TrimPrefix(fileName, "/")))

	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(resp.Status())
	}

	return resp.String(), nil
}

func (c *FileClient) PublishFile(meshName, fileName string, content []byte) error {
	resp, err := c.httpClient.R().
		SetBody(content).
		Post(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "file-data", strings.TrimPrefix(fileName, "/")))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *FileClient) EraseFile(meshName, fileName string) error {
	resp, err := c.httpClient.R().
		Delete(fmt.Sprintf("%s/%s/%s/%s", "meshes", meshName, "file-data", strings.TrimPrefix(fileName, "/")))

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusNoContent && resp.StatusCode() != http.StatusNotFound {
		return errors.New(resp.Status())
	}

	return nil
}
