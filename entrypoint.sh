#!/bin/bash

# 后台启动 GRPC
./grpc_server &
# 前台启动 HTTP
./http_server
