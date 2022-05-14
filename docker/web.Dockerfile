FROM golang:1.18-buster

ENV GO118MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://mirrors.aliyun.com/goproxy/,direct" \
    GOPRIVATE="*github.com" \
    GOSUMDB=off \
    GOPATH="/GODIR/"

RUN go version
RUN mkdir /GODIR
WORKDIR /build
COPY .. .
WORKDIR /build
RUN rm /build/config.development.yml
RUN cp /build/config/config.production.yml /build/config.yml
RUN go build -o main .
EXPOSE 8081
CMD ["/build/main"]