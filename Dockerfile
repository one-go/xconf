FROM golang:1.14 AS builder

ENV CGO_ENABLED=0

COPY . /go/src/github.com/one-go/xconf

RUN cd /go/src/github.com/one-go/xconf && make release-binary

FROM scratch
COPY --from=builder bin/xconf /usr/bin/xconf
COPY --from=builder web /web

ENTRYPOINT ["/usr/bin/xconf"]
