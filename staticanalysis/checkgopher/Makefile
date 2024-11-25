build:
	GOOS=darwin GOARCH=amd64 go build -o checkgopher cmd/checkgopher/main.go

MYVETTOOL := $(shell pwd)/checkgopher

myvet:
	go vet -vettool="$(MYVETTOOL)" ./...
