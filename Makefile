.PHONY: build clean install

VERSION=$(shell git describe --tags --dirty)
BINARY=bin/jt-$(shell go env GOOS)-$(shell go env GOARCH)

build: $(BINARY)

version:
	@echo version $(VERSION)

$(BINARY): *.go
	@echo building $(VERSION)
	@mkdir -p bin
	@go build -o $(BINARY) -ldflags "-X main.version $(VERSION)"

clean:
	@echo cleaning
	@rm -rf bin

prefix=/usr/bin
install: $(BINARY)
	@echo installing $(BINARY) to $(prefix)/jt
	@install -D $(BINARY) $(prefix)/jt
