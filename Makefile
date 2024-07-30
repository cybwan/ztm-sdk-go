#!make

CTR_REGISTRY ?= flomesh
CTR_TAG      ?= latest

DOCKER_BUILDX_PLATFORM ?= linux/amd64
DOCKER_BUILDX_OUTPUT ?= type=registry

.PHONY: buildx-context
buildx-context:
	@if ! docker buildx ls | grep -q "^ztm "; then docker buildx create --name ztm --driver-opt network=host; fi

.PHONY: docker-build-ztm-hub
docker-build-ztm-hub:
	docker buildx build --builder fsm --platform=$(DOCKER_BUILDX_PLATFORM) -o $(DOCKER_BUILDX_OUTPUT) -t $(CTR_REGISTRY)/ztm-hub:$(CTR_TAG) -f dockerfiles/Dockerfile.hub .

.PHONY: docker-build-ztm-agent
docker-build-ztm-agent:
	docker buildx build --builder fsm --platform=$(DOCKER_BUILDX_PLATFORM) -o $(DOCKER_BUILDX_OUTPUT) -t $(CTR_REGISTRY)/ztm-agent:$(CTR_TAG) -f dockerfiles/Dockerfile.agent .

ZTM_TARGETS = ztm-hub ztm-agent

DOCKER_ZTM_TARGETS = $(addprefix docker-build-, $(ZTM_TARGETS))

$(foreach target,$(ZTM_TARGETS),$(eval docker-build-$(target): buildx-context))

.PHONY: docker-build-ztm
docker-build-ztm: $(DOCKER_ZTM_TARGETS)

.PHONY: docker-build-cross
docker-build-cross: DOCKER_BUILDX_PLATFORM=linux/amd64,linux/arm64
docker-build-cross: docker-build-ztm

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

.PHONY: TestPublishFile
TestPublishFile:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestPublishFile

.PHONY: TestDescribeFile
TestDescribeFile:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestDescribeFile

.PHONY: TestDownloadFile
TestDownloadFile:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestDownloadFile

.PHONY: TestEraseFile
TestEraseFile:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestEraseFile

.PHONY: TestListFiles
TestListFiles:
	CTR_AGENT=$(CTR_AGENT) go test -run=TestListFiles