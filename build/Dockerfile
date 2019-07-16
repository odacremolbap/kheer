ARG GO_IMAGE=golang:1.12.7

FROM $GO_IMAGE AS build
WORKDIR /kheer

COPY go.mod .
COPY go.sum .

RUN go mod download
