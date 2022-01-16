# build stage
FROM golang:alpine AS build-env
RUN apk add --no-cache git make
COPY . /go/src/whatismyip
RUN GO111MODULE=on
RUN cd /go/src/whatismyip && make build-linux

# final stage
FROM alpine
WORKDIR /whatismyip
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*
COPY --from=build-env /go/src/whatismyip/bin/whatismyip-linux-amd64 ./whatismyip
COPY --from=build-env /go/src/whatismyip/web ./web
CMD ./whatismyip
