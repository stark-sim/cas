FROM --platform=$BUILDPLATFORM golang:1.19-alpine AS builder

LABEL maintainer="StarkSim<gooda159753@163.com>"

# 在容器根目录创建 src 目录
WORKDIR /src

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache g++

COPY ./go.mod .

COPY ./go.sum .

ENV GOPROXY="https://goproxy.cn"

RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
# 由于获取 config.yaml 位置时需要依赖文件内部路径，所以不能加 -trimpath
RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-s -w" -o http_server ./internal/cas_http/main.go

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/http_server /app/
COPY --from=builder /src/internal/db/migrations /app/internal/db/migrations/

EXPOSE 8080

ENTRYPOINT ["./http_server"]
