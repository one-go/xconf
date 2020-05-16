FROM golang:1.14 AS builder

ENV CGO_ENABLED=0

COPY . /go/src/github.com/one-go/xconf

RUN cd /go/src/github.com/one-go/xconf && make

FROM scratch
COPY --from=builder /go/src/github.com/one-go/xconf/bin/xconf /usr/bin/xconf
COPY --from=builder /go/src/github.com/one-go/xconf/web /web

ENTRYPOINT ["xconf"]
