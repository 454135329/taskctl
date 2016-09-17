VERSION = $(shell git describe $(shell git rev-list --tags --max-count=1))
BINARY = taskctl

deps:
	go get -u github.com/olekukonko/tablewriter
	go get -u github.com/urfave/cli

build: deps
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/$(BINARY) ./

install: build
	mv bin/$(BINARY) $(GOPATH)/bin

# Use this command in subshel: $(make run) list
run:
	@echo go run ./taskctl.go
