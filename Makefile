.PHONY: clean build install

default: build

build: bin/jt

bin/jt: *.go
	@mkdir -p bin
	@go build -o bin/jt

clean:
	@rm -rf bin

install: bin/jt
	@sudo cp bin/jt /usr/bin/jt
	@sudo chown root /usr/bin/jt
