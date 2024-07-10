FROM golang:1.22.5

WORKDIR /go/src/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
