FROM alpine:3.21
RUN apk update && apk add git

USER root

COPY . /go-authenticator
WORKDIR /go-authenticator

RUN go mod tidy
RUN go build -o /go-sqlserver github.com/cybergarage/go-sqlserver/cmd/go-sqlserver

COPY ./sql/conf/go-sqlserver.yaml /
COPY ./docker/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]
