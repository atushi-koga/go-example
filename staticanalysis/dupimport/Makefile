build:
	GOOS=darwin GOARCH=amd64 go build -o dupimport cmd/dupimport/main.go

MYVETTOOL := $(shell pwd)/dupimport

myvet:
	go vet -vettool="$(MYVETTOOL)" ./...
