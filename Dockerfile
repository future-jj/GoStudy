# Dockerfile
FROM golang:1.22.4 AS builder

WORKDIR /app
COPY . .

# 构建 Go 程序
RUN go build -o main .

# 运行阶段
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
