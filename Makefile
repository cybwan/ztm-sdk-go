#!make

CTR_AGENT ?= home
MY_HOST_IP ?= 127.0.0.1

.PHONY: go-test
go-test:
	go test -v ./...; echo $?

.PHONY: TestJoin
TestJoin:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestJoin

.PHONY: TestListEndpoints
TestListEndpoints:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestListEndpoints

.PHONY: TestStartApp
TestStartApp:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestStartApp

.PHONY: TestOpenOutbound
TestOpenOutbound:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestOpenOutbound

.PHONY: TestOpenInbound
TestOpenInbound:
	CTR_AGENT=$(CTR_AGENT) MY_HOST_IP=$(MY_HOST_IP) go test -run=TestOpenInbound