ARG NAMESPACE
FROM ${NAMESPACE}/golang-base AS build
WORKDIR /go/

COPY src src
WORKDIR /go/src/components/notification-server
RUN go build -o notification-server


FROM ${NAMESPACE}/golang-base
WORKDIR /app/
COPY --from=build /go/src/components/notification-server/notification-server /app/
CMD ["./notification-server"]


