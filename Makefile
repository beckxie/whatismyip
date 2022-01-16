APP_NAME=whatismyip

OS_LINUX=linux
OS_DARWIN=darwin
OS_WINDOWS=windows

GO_BUILD=go build -o bin/$(APP_NAME)

GO_BUILD_LINUX=GOOS=$(OS_LINUX) CGO_ENABLED=0 GOARCH=amd64 && $(GO_BUILD)-$(OS_LINUX)-amd64
GO_BUILD_DARWIN=GOOS=$(OS_DARWIN) CGO_ENABLED=0 GOARCH=amd64 && $(GO_BUILD)-$(OS_DARWIN)-amd64
GO_BUILD_WINDOWS=GOOS=$(OS_WINDOWS) CGO_ENABLED=0 GOARCH=amd64 && $(GO_BUILD)-$(OS_WINDOWS)-amd64.exe

DOCKER_TAG?=latest
DOCKER_OWNER?=beckxie
DOCKER_BUILD_DESC=$(info Make: Building "$(DOCKER_OWNER)/$(APP_NAME):$(DOCKER_TAG)" images.)
DOCKER_BUILD=docker build -t $(DOCKER_OWNER)/$(APP_NAME):$(DOCKER_TAG) .
DOCKER_PUSH_DESC=$(info Make: Pushing "$(DOCKER_OWNER)/$(APP_NAME):$(DOCKER_TAG)" images.)
DOCKER_PUSH=docker push $(DOCKER_OWNER)/$(APP_NAME):$(DOCKER_TAG)

.PHONY: all
all: lint test build

.PHONY: build
build: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux: 
	$(GO_BUILD_LINUX)

.PHONY: build-darwin
build-darwin: 
	$(GO_BUILD_DARWIN)

.PHONY: build-windows
build-windows: 
	$(GO_BUILD_WINDOWS)

.PHONY: test
test:
	go test -cover -count=1 ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: docker-build
docker-build: Dockerfile
	$(DOCKER_BUILD_DESC)
	$(DOCKER_BUILD)

.PHONY: docker-push
docker-push:
	$(DOCKER_PUSH_DESC)
	$(DOCKER_PUSH)
