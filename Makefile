GOFMT ?= gofmt "-s"
PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

all: install

install:
	go get github.com/gonitor/gonitor
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	sh coverage.sh

build:
	sh BuildMulti.sh

buildDocker:
	docker build -t gonitor/gonitor .