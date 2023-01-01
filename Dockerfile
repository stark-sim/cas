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

RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags "-s -w" -o http_server ./internal/cas_http/main.go
RUN CGO_ENABLE=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags "-s -w" -o grpc_server ./internal/cas_grpc/main.go

FROM alpine:latest

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /src/http_server /app/
COPY --from=builder /src/grpc_server /app/
COPY --from=builder /src/internal/db/migrations /app/internal/db/migrations/

EXPOSE 8080, 8081

ENTRYPOINT ["./http_server", "./grpc_server"]
