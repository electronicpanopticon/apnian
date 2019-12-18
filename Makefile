# Basic go commands
GOCMD=go
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test

.PHONY: perf release test tidy

all: test

perf:
		$(GOTEST) -bench=Fact20 -cpuprofile=cpu.out
		$(GOCMD) tool pprof -png  apnian.go.test cpu.out

release: tidy
		$(GOCMD) fmt

test:
		$(GOTEST) -v ./... -cover

tidy:
		$(GOMOD) tidy
		$(GOCMD) list -m all
