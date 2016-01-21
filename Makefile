default: test

export GOPATH=$(CURDIR)/
export GOBIN=$(CURDIR)/.temp/

build:
	go build .

test: build
	go test
	go test -bench=.

fmt:
	@go fmt .
