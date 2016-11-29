all: build test

build:
	go build

deps:
	go get -u github.com/jinzhu/gorm

release: build
	strip eir

test:
	go test

clean:
	rm eir

.PHONY: all build deps release clean test