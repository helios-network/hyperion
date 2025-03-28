APP_VERSION = $(shell git describe --abbrev=0 --tags)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date -u "+%Y%m%d-%H%M")
VERSION_PKG = github.com/Helios-Chain-Labs/hyperion/orchestrator/version
IMAGE_NAME := gcr.io/helios-core/hyperion
DOCKER ?= false

all:

image:
	docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local -f Dockerfile .
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push:
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):latest

install: export GOPROXY=direct
install: export VERSION_FLAGS="-X $(VERSION_PKG).GitCommit=$(GIT_COMMIT) -X $(VERSION_PKG).BuildDate=$(BUILD_DATE)"
install:
	@$(DOCKER) && go install -tags muslc -ldflags $(VERSION_FLAGS) ./cmd/... || go install -ldflags $(VERSION_FLAGS) ./cmd/...
	@echo "hyperion command line installed in your go path"

.PHONY: install image push test gen

start:
	hyperion orchestrator

test:
	# go clean -testcache
	go test ./test/...
