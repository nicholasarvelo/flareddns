# Project metadata
IMAGE_NAME = flareddns
VERSION    ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
COMMIT     = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
REPO_URL   = https://github.com/steptimeeditor/flareddns

.PHONY: all build push

all: build

build:
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		-t $(IMAGE_NAME):$(VERSION) .

push:
	docker push $(IMAGE_NAME):$(VERSION)
