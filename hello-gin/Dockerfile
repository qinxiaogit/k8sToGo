# 多阶段构建：提升构建速度，减少镜像大小

# 从官方仓库中获取 1.20 的 Go 基础镜像
# syntax=docker/dockerfile:1
FROM --platform=$TARGETPLATFORM golang:1.20-alpine AS builder
ARG TARGETARCH
ARG TARGETOS

WORKDIR  /workspace

# 安装项目依赖
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# 复制项目文件
COPY . .

# 构建名为"app"的二进制文件
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o app main.go

# 获取 Distroless 镜像，只有 650 kB 的大小，是常用的 alpine:latest 的 1/4
FROM gcr.io/distroless/static:nonroot-$TARGETARCH
# 设置工作目录
WORKDIR /
# 将上一阶段构建好的二进制文件复制到本阶段中
COPY --from=builder /workspace/app .
# 设置监听端口
EXPOSE 8080
# 配置启动命令
ENTRYPOINT ["/app"]
