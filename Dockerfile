# 第一阶段：构建阶段
FROM golang:1.23.3-alpine3.20 AS builder

LABEL maintainer="AptS:1547 <esaps@esaps.net>"

WORKDIR /app

# 复制 Go 模块文件
COPY go.mod go.sum /app/

# 设置 Go 模块代理
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download

# 复制项目代码
ADD . /app

# 编译项目，生成可执行文件 ad
RUN go build -o ad

# 第二阶段：运行阶段
FROM ubuntu:22.04


WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/ad /app/ad

# 从构建阶段复制配置文件（如果需要）
COPY --from=builder /app/.env /app/.env

# 从构建阶段复制模板文件（如果需要）
COPY --from=builder /app/templates /app/templates

# 设置容器启动命令
CMD ["./ad"]