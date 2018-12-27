TAG_NAME=$(shell git describe --abbrev=0 --tags)
BUILD_NAME=nie-api
MAC_DIST=dist/$(BUILD_NAME)-darwin-amd64-$(TAG_NAME).app
LIN_DIST=dist/$(BUILD_NAME)-linux-amd64-$(TAG_NAME)
WIN_DIST=dist/$(BUILD_NAME)-windows-amd64-$(TAG_NAME).exe

get:
	go get -u github.com/gin-gonic/gin/...
	go get -u github.com/humbertodias/go-nie-crawler

run:
	go run main.go

init:
	mkdir -p dist

build-mac:  init
	GOOS=darwin GOARCH=amd64 go build -o $(MAC_DIST)
	chmod +x $(MAC_DIST)

build-lin:  init
	GOOS=linux GOARCH=amd64 go build -o $(LIN_DIST)
	chmod +x $(LIN_DIST)

build-win:  init
	GOOS=windows GOARCH=amd64 go build -o $(WIN_DIST)
	chmod +x $(WIN_DIST)

build: build-mac    build-lin   build-win

test:
	go test -v

lint:
	golint

format:
	go fmt

fix:
	go fix

clean-dist:
	rm -rf dist

clean:  clean-dist
