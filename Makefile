VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

get:
	go get

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build: get: format
	go build -v -o lawBot -ldflags "-X="github.com/pancheliuga/law-bot/cmd.appVersion=${VERSION}

clean:
	rm -rf lawBot