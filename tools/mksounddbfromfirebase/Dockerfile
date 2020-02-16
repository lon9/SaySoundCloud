FROM golang:alpine as build-env

ARG GITHUB_ACCESS_TOKEN
RUN apk add --no-cache git gcc musl-dev
RUN git config --global url."https://${GITHUB_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/lon9/SaySoundCloud/tools/mksounddbfromfirebase

RUN go get -v github.com/lon9/SaySoundCloud/tools/mksounddbfromfirebase
RUN go build -o /usr/bin/app

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=build-env /usr/bin/app /usr/bin/app
