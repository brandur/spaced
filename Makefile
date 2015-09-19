SHELL = /bin/sh

all: test

build:
	go build ./...

clean:
	go clean ./...

fmt:
	gofmt -w .

install:
	go install ./...

run: forego install
	forego start

test:
	go test ./...

update-deps: godep
	godep save -r ./...

# Dependencies

forego:
	go get github.com/ddollar/forego

godep:
	go get github.com/tools/godep
