# Basic go commands
GOCMD=go
GOTEST=$(GOCMD) test

.PHONY: test

all: get test

test:
		$(GOTEST) -v ./...

get:
		dep ensure
