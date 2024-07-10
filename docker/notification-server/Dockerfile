ARG NAMESPACE
FROM ${NAMESPACE}/golang-base
WORKDIR /go/src/

COPY notification-server notification-server
WORKDIR /go/src/notification-server
RUN go build -o notification-server

CMD ["./notification-server"]


