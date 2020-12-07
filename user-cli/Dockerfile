FROM golang:1.14-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk --no-cache add git

WORKDIR /app/laracom-user-cli

COPY . .

RUN go mod download && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-user-cli

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk update && apk add --no-cache bash ca-certificates && \
mkdir /app
WORKDIR /app

COPY --from=builder /app/laracom-user-cli/laracom-user-cli .

ENTRYPOINT ["./laracom-user-cli"]