PROJ=xconf
VERSION=`git rev-parse`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o bin/xconf
proto:
	go generate github.com/one-go/xconf/...
