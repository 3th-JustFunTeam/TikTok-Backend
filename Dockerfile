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
COPY . .
WORKDIR /build
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .
EXPOSE 8081
CMD ["/dist/main"]