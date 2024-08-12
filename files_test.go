package ztm_go_sdk

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	user = "root"
	file = "xxx"
)

func TestListFiles(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	files, _ := client.ListFiles(CurrCtx.meshName)
	bytes, _ := json.MarshalIndent(files, "", " ")
	fmt.Println(string(bytes))
}

func TestPublishFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	content := client.PublishFile(CurrCtx.meshName, fmt.Sprintf("%s/%s/%s", BaseFolder, user, file), []byte("demo"))
	fmt.Println(content)
}

func TestDescribeFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	file, _ := client.DescribeFile(CurrCtx.meshName, fmt.Sprintf("%s/%s/%s", BaseFolder, user, file))
	bytes, _ := json.MarshalIndent(file, "", " ")
	fmt.Println(string(bytes))
}

func TestDownloadFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	content, _ := client.DownloadFile(CurrCtx.meshName, fmt.Sprintf("%s/%s/%s", BaseFolder, user, file))
	fmt.Println(content)
}

func TestEraseFile(t *testing.T) {
	client := FileClient{
		RestClient: NewRestClient(CurrCtx.agentAddr, false),
	}
	client.httpClient.Debug = true
	err := client.EraseFile(CurrCtx.meshName, fmt.Sprintf("%s/%s/%s", BaseFolder, user, file))
	fmt.Println(err)
}
