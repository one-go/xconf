ARG GO_VERSION=1.14
FROM golang:${GO_VERSION}-alpine as build

# Necessary to run 'go get' and to compile the linked binary
RUN apk add git musl-dev

ADD . /xconf

WORKDIR /xconf

ENV GO111MODULE=on

# build & install server
RUN go get -u ./... && make build

FROM scratch AS final
LABEL maintainer="ifish <fishioon@gmail.com>"

COPY --from=build /xconf/bin/xconf /go/bin/xconf

ENTRYPOINT ["/go/bin/xconf", "--listener", ":8900"]

EXPOSE 8900
