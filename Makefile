# Basic go commands
GOCMD=go
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test

.PHONY: release test tidy

all: test

release: tidy
		$(GOCMD) fmt

test:
		$(GOTEST) -v ./... -cover

tidy:
		$(GOMOD) tidy
		$(GOCMD) list -m all
