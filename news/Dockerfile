FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn"

WORKDIR /build

COPY . .

RUN go build -o main ./cmd/.

WORKDIR /dist

RUN cp /build/main .

EXPOSE all

CMD ["/dist/main"]
