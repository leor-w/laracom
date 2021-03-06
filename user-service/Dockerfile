# user-service/Dockerfile

# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:1.14-alpine as builder

# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn, direct
ENV TIME_ZONE Asia/Shanghai

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk --no-cache add git && \
ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && \
echo $TIME_ZONE > /etc/timezone

# 设置工作目录
WORKDIR /app/laracom-user-service

# 将当前项目所在目录代码拷贝到镜像中
COPY . .

# 下载依赖
RUN go mod tidy && \
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o laracom-user-service

# 下面是第二阶段的镜像构建，和之前保持一致
FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk update && apk add --no-cache bash ca-certificates && \
mkdir /app
WORKDIR /app

# 这一步不再从宿主机拷贝二进制文件，而是从上一个阶段构建的 builder 容器中拉取，神奇吧
COPY --from=builder /app/laracom-user-service/laracom-user-service .

# 和之前一样通过 Supervisor 管理服务
CMD ["./laracom-user-service"]