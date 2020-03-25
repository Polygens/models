FROM golang:1.13-alpine AS builder
RUN apk update && apk add --no-cache git
RUN go get golang.org/x/tools/cmd/godoc
RUN rm -rf /go/src/*
WORKDIR /go/src/github.com/polygens/models
RUN go get github.com/paulmach/go.geojson
COPY . .
RUN echo "appuser:x:65534:65534:appuser:/:" > /etc_passwd

FROM alpine
COPY --from=builder /go/bin/godoc /
COPY --from=builder /go/src /documentation/src
COPY --from=builder /etc_passwd /etc/passwd
USER appuser
WORKDIR /documentation/src/github.com/polygens/models
ENV GOROOT=/documentation
ENTRYPOINT ["/godoc", "-http=:6060", "-index", "-index_interval=-1m"]
