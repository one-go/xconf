FROM golang:1.14 AS builder

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/one-go/xconf
COPY . .

RUN go build -ldflags "-X main.Version=`git rev-parse --short HEAD` -X main.Build=`date +%FT%T%z`" -o /go/bin/xconf github.com/one-go/xconf/console

FROM scratch
COPY --from=builder /go/bin/xconf /usr/bin/xconf

ENTRYPOINT ["/usr/bin/xconf"]
