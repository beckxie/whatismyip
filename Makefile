.PHONY: all
all: whereismyip-linux-amd64 whereismyip-windows-amd64

.PHONY: whereismyip-linux-amd64
whereismyip-linux-amd64:
	@echo "Building linux binary..."
	@CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -mod=vendor -o bin/whereismyip_linux_amd64

.PHONY: whereismyip-windows-amd64
whereismyip-windows-amd64:
	@echo "Building windows binary..."
	@CGO_ENABLED=0 \
	GOOS=windows \
	GOARCH=amd64 \
	go build -mod=vendor -o bin/whereismyip_windows_amd64