ARG NAMESPACE
FROM ${NAMESPACE}/golang-base
WORKDIR /go/src/

COPY components components
WORKDIR /go/src/components/notification-server
RUN go build -o notification-server

CMD ["./notification-server"]


