all: build lint fmt 

build:
	go build ./...

fmt:
	go fmt ./...

lint:
	golint ./...

