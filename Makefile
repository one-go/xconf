PROJ=xconf
BUILD=`date +%FT%T%z`

VERSION=`git rev-parse --short HEAD`
TAG=$(git describe --exact-match --abbrev=0 --tags ${COMMIT} 2> /dev/null || true)
[ -z "$TAG" ] || VERSION=$TAG
[ -n "$(git diff --shortstat 2> /dev/null | tail -n1)" ] && VERSION="${VERSION}-dirty"

DOCKER_REPO=onego/xconf
DOCKER_IMAGE=$(DOCKER_REPO):$(VERSION)

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o bin/xconf
proto:
	go generate github.com/one-go/xconf/...

.PHONY: docker-image
docker-image:
	docker build -t $(DOCKER_IMAGE) .
