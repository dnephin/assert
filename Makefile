
BUILD_ID ?= $(USER)
DOCKER_IMAGE := assert-dev:$(BUILD_ID)
PACKAGE := github.com/dnephin/assert
SRC := /go/src/$(PACKAGE)

VOLUMES := \
	-v $(CURDIR):$(SRC) \
	-v $(CURDIR)/dist/bin:/go/bin \
	-v $(CURDIR)/dist/pkg:/go/pkg


all: test



build: .dev-image

.dev-image: Dockerfile.build
	docker build -t $(DOCKER_IMAGE) -f Dockerfile.build .
	touch $@

Dockerfile.build:



dist:
	mkdir dist/

.PHONY: shell
shell: dist build
	@docker run --rm -ti $(VOLUMES) $(DOCKER_IMAGE) bash

test-unit: build
	docker run --rm -ti $(VOLUMES) $(DOCKER_IMAGE) go test -v ./...

test: test-unit
