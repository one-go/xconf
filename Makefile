REPO_PATH=github.com/one-go/xconf
BUILD=`date +%FT%T%z`

VERSION ?= $(shell ./scripts/git-version)

DOCKER_REPO=onego/xconf
DOCKER_IMAGE=$(DOCKER_REPO):$(VERSION)

LD_FLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LD_FLAGS} -o bin/xconf ${REPO_PATH}/console
proto:
	protoc -I api --go_out=plugins=grpc:api --js_out=import_style=commonjs:web --grpc-web_out=import_style=commonjs,mode=grpcwebtext:web api/xconf.proto

.PHONY: docker-image
docker-image:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: web
web:
	cd web && npx webpack index.js
