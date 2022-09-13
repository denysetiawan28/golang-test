FROM golang:1.19-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git && \
    apk add --no-cache tzdata && \
    apk add --no-cache curl

MAINTAINER denysetiawan28

ADD . /opt/apps
WORKDIR /opt/apps

COPY ./resources /opt/apps/resources

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o echo-test

FROM alpine:3.16

COPY --from=builder /opt/apps/echo-test /opt/apps/echo-test
COPY --from=builder /opt/apps/resources /opt/apps/resources

EXPOSE 1234

CMD ["/opt/apps/echo-test"]
