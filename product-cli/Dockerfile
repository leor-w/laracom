FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn, direct
ENV TIME_ZONE Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk --no-cache add git

WORKDIR /app/laracom-product-cli

COPY . .

RUN go mod tidy && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-product-cli

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk update && apk add --no-cache bash ca-certificates && \
mkdir /app
WORKDIR /app

COPY --from=builder /app/laracom-product-cli/laracom-product-cli .

ENTRYPOINT ["./laracom-product-cli"]