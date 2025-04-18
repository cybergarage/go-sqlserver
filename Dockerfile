FROM alpine:latest
RUN apk update && apk add git go

USER root

COPY . /go-authenticator
WORKDIR /go-authenticator

RUN go mod tidy
RUN go build -o /go-sqlserver github.com/cybergarage/go-sqlserver/cmd/go-sqlserver

COPY ./sql/conf/go-sqlserver.yaml /
COPY ./docker/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]
