FROM golang:alpine as build-env

RUN apk add --no-cache git

ADD . /go/src/github.com/lon9/SaySoundCloud/backend
WORKDIR /go/src/github.com/lon9/SaySoundCloud/backend

RUN go get -v && \
    go build -o /usr/bin/app

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=build-env /usr/bin/app /usr/bin/app
