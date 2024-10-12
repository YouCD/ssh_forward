BINARY_DIR=bin/ssh_forward
BINARY_NAME=ssh_forward
check:
	@golangci-lint run ./...
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BINARY_DIR)/$(BINARY_NAME)