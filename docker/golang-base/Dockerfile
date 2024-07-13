FROM golang:1.22.5-alpine3.19

WORKDIR /go/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
