#/bin/bash

mkdir -p out
go test -v -vet=all ./... -covermode=count -coverprofile=out/coverage.out
