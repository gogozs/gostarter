FROM golang:latest

MAINTAINER zs "demo@qq.com"

WORKDIR /app
ENV GOPROXY  https://gocenter.io

ADD . /app
RUN go build  -mod=vendor  main.go

ENTRYPOINT ["./main"]
