### God - Godly Golang Tools ###

VERSION := $(shell git describe --tags --always --dirty) # Current git tag/hash.

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

version:
	@echo $(VERSION)

.DEFAULT_GOAL := help
MAKEFLAGS += --no-print-directory # Don't print unnecessary output.

#-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~#

all:
	@'$(MAKE)' test

test:
	go mod tidy
	go test ./... -race -cover

clean:
	go clean -cache -modcache -testcache

push:
	git add .
	git commit -m "[GOD] - $(msg)"
	git push origin master

