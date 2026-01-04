APP_NAME=whatismyip
BIN_DIR=bin
CMD_PATH=./cmd/whatismyip

OS_LINUX=linux
OS_DARWIN=darwin
OS_WINDOWS=windows

DOCKER_TAG?=latest
DOCKER_OWNER?=beckxie
DOCKER_IMAGE=$(DOCKER_OWNER)/$(APP_NAME):$(DOCKER_TAG)

.PHONY: all
all: lint test build

.PHONY: run
run:
	go run $(CMD_PATH)

.PHONY: build
build: build-linux build-darwin build-windows

.PHONY: build-local
build-local:
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

.PHONY: build-linux
build-linux:
	GOOS=$(OS_LINUX) GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-$(OS_LINUX)-amd64 $(CMD_PATH)

.PHONY: build-darwin
build-darwin:
	GOOS=$(OS_DARWIN) GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-$(OS_DARWIN)-amd64 $(CMD_PATH)
	GOOS=$(OS_DARWIN) GOARCH=arm64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-$(OS_DARWIN)-arm64 $(CMD_PATH)

.PHONY: build-windows
build-windows:
	GOOS=$(OS_WINDOWS) GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(APP_NAME)-$(OS_WINDOWS)-amd64.exe $(CMD_PATH)

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: lint
lint:
	go vet ./...
	@which golangci-lint > /dev/null && golangci-lint run || echo "golangci-lint not installed, skipping"

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_IMAGE)

.PHONY: docker-run
docker-run:
	docker run --rm -p 9999:9999 $(DOCKER_IMAGE)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)/*
