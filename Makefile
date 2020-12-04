.PHONY: default clean install lint test assets build binaries test-release release publish-testing publish-latest publish-images

TAG_NAME := 2020.0.34
SHA := $(shell test -d .git && git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

IMAGE := mark-sch/evcc
ALPINE := 3.12
TARGETS := arm.v6,arm.v8,amd64

default: clean install npm assets lint test build

clean:
	rm -rf dist/

install:
	go install github.com/mjibson/esc
	go install github.com/golang/mock/mockgen
	npm ci

lint:
	golangci-lint run

test:
	@echo "Running testsuite"
	go test ./...

npm:
	npm run build

ui:
	npm run build
	go generate main.go

assets:
	@echo "Generating embedded assets"
	go generate ./...

build:
	@echo Version: $(VERSION) $(BUILD_DATE)
	go build -v -tags=release -ldflags '-X "github.com/mark-sch/evcc/server.Version=${VERSION}" -X "github.com/mark-sch/evcc/server.Commit=${SHA}"'

release-test:
	goreleaser --snapshot --skip-publish --rm-dist

release:
	goreleaser --rm-dist

publish-testing:
	@echo Version: $(VERSION) $(BUILD_DATE)
	seihon publish --dry-run=false --template docker/tmpl.Dockerfile --base-runtime-image alpine:$(ALPINE) \
	   --image-name $(IMAGE) -v "testing" --targets=arm.v6,amd64

publish-latest:
	@echo Version: $(VERSION) $(BUILD_DATE)
	seihon publish --dry-run=false --template docker/tmpl.Dockerfile --base-runtime-image alpine:$(ALPINE) \
	   --image-name $(IMAGE) -v "latest" --targets=$(TARGETS)

publish-images:
	@echo Version: $(VERSION) $(BUILD_DATE)
	seihon publish --dry-run=false --template docker/tmpl.Dockerfile --base-runtime-image alpine:$(ALPINE) \
	   --image-name $(IMAGE) -v "latest" -v "$(TAG_NAME)" --targets=$(TARGETS)
