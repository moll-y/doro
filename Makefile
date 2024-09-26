GOCMD=go
GOFMT=gofmt
GOFILES=$(shell find . -name '*.go')

.PHONY: fmt

fmt:
	$(GOFMT) -w -s $(GOFILES)
