.PHONY: all
all: whatismyip-linux-amd64 whatismyip-windows-amd64

.PHONY: whatismyip-linux-amd64
whatismyip-linux-amd64:
	@echo "Building linux binary..."
	@CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -mod=vendor -o bin/whatismyip_linux_amd64

.PHONY: whatismyip-windows-amd64
whatismyip-windows-amd64:
	@echo "Building windows binary..."
	@CGO_ENABLED=0 \
	GOOS=windows \
	GOARCH=amd64 \
	go build -mod=vendor -o bin/whatismyip_windows_amd64