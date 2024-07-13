ARG NAMESPACE
FROM ${NAMESPACE}/golang-base AS build
WORKDIR /go/

COPY src src
WORKDIR /go/src/components/notification-server
RUN go build -o notification-server

CMD ["./notification-server"]

# FROM golang:1.22.5-alpine3.19
# WORKDIR /app/
# COPY --from=build /go/src/components/notification-server/notification-server /app/
# CMD ["./notification-server"]
# CMD ["sh"]


