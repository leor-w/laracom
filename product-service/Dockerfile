# product-service/Dockerfile
FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io, direct
ENV TIME_ZONE Asia/Shanghai

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk update && \
apk add --no-cache git && \
ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && \
echo $TIME_ZONE > /etc/timezone

WORKDIR /app/laracom-product-service

COPY . .

RUN go mod tidy && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-product-service

FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk update && \
apk add --no-cache bash ca-certificates && \
mkdir /app

WORKDIR /app

COPY --from=builder /app/laracom-product-service/laracom-product-service .

CMD ["./laracom-product-service"]
