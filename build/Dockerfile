# This image is used when compiling (see Makefile)
# It is created from a Go image, and added dependencies for caching

# GO_IMAGE environment can be set to override base image
ARG GO_IMAGE=golang:1.12.7

FROM $GO_IMAGE AS build
WORKDIR /kheer

COPY go.mod .
COPY go.sum .

RUN go mod download
