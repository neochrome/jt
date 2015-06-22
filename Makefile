.PHONY: clean build

default: build

build: bin/jt

bin/jt: *.go
	@mkdir -p bin
	@go build -o bin/jt

clean:
	@rm -rf bin
