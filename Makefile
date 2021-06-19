.PHONY: build all

build:
	go build -o kaiji-admin

build-linux:
	GOOS=linux GOARCH=amd64 go build -o kaiji-admin
