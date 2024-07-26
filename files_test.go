package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestListFiles(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	files, _ := client.ListFiles(currCtx.meshName)
	bytes, _ := json.MarshalIndent(files, "", " ")
	fmt.Println(string(bytes))
}

func TestPublishFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	content := client.PublishFile(currCtx.meshName, "/home/root/xxx", []byte("demo"))
	fmt.Println(content)
}

func TestDescribeFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	file, _ := client.DescribeFile(currCtx.meshName, "/home/root/xxx")
	bytes, _ := json.MarshalIndent(file, "", " ")
	fmt.Println(string(bytes))
}

func TestDownloadFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	content, _ := client.DownloadFile(currCtx.meshName, "/home/root/xxx")
	fmt.Println(content)
}

func TestEraseFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(currCtx.agentAddr),
	}
	err := client.EraseFile(currCtx.meshName, "/home/root/xxx")
	fmt.Println(err)
}
