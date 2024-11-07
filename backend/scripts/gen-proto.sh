#!/bin/bash

# 确保在项目根目录下执行
cd "$(dirname "$0")/.."

# 创建输出目录
mkdir -p pb

# 编译 proto 文件
protoc --proto_path=proto \
    --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/file_service.proto