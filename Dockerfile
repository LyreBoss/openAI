# 使用 golang 官方镜像作为基础镜像
FROM golang:1.17-alpine

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到容器的 /app 目录
COPY . .

# 构建 Go 项目
RUN go build -o main .

EXPOSE 8080

# 设置容器启动时执行的命令
CMD ["./cmd/main"]
