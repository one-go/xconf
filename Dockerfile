FROM golang:1.14 AS builder

ENV CGO_ENABLED=0

WORKDIR /go/src/xconf
COPY . .

RUN go install -v -ldflags "-X main.Version=`git rev-parse HEAD` -X main.Build=`date +%FT%T%z`"

FROM scratch
COPY --from=builder /go/bin/xconf /usr/bin/xconf

ENTRYPOINT ["/usr/bin/xconf"]
