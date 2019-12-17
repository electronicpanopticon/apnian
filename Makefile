# Basic go commands
GOCMD=go
GOTEST=$(GOCMD) test

.PHONY: test

all: test

test:
		$(GOTEST) -v ./... -cover
