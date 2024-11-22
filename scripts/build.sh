#/bin/bash

go mod tidy && \
go fmt ./... && \
go vet ./... && \
go fix ./... && \
gosec ./... && \
staticcheck -go 1.23.3 ./...
go build -v ./...
