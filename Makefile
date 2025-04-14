.DEFAULT_GOAL := build

.PHONY: fmt vet build
fmt:
	go fmt ./..

vet: fmt
	go vet ./..

build: fmt vet
	go build ./..
	@echo "Build complete."