# 使用 Go 的官方镜像作为基础
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制源码
COPY . .

# 编译应用
RUN go build -o main .

# 使用轻量级的镜像
FROM alpine:latest

# 复制编译后的二进制文件
COPY --from=builder /app/main .

# 指定运行的命令
CMD ["./main"]
