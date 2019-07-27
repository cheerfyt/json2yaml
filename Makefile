##### json2yaml

all: build

build:
	@go build -o json2yaml main.go

install: build
	@cp ./json2yaml /usr/local/bin/

clean:
	@rm ./json2yaml

.PHONY: build install build-image clean
