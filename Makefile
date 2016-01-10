.PHONY: .build clean install
.DEFAULT_GOAL := all

VERSION=$(shell git describe --tags --dirty 2>/dev/null || echo 'dev')
NAME=jt

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BINARY=bin/$(NAME)-$(GOOS)-$(GOARCH)

bin:
	@mkdir -p ./bin

build: $(BINARY)
$(BINARY): *.go .build-auto

.build%: bin
	@echo building version: $(VERSION) for $(GOOS)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o ./bin/$(NAME)-$(GOOS)-$(GOARCH)$(EXT) -ldflags "-X main.version=$(VERSION)"

all: linux win osx

linux: GOOS = linux
linux: GOARCH = amd64
linux: .build-linux

win: GOOS = windows
win: GOARCH = amd64
win: EXT = .exe
win: .build-win

osx: GOOS = darwin
osx: GOARCH = amd64
osx: .build-osx

clean:
	@echo cleaning
	@rm -rf bin

prefix=/usr/bin
install: $(BINARY)
	@echo installing $(BINARY) to $(prefix)/$(NAME)
	@install -D $(BINARY) $(prefix)/$(NAME)
