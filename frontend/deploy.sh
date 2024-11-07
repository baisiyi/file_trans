#!/bin/bash

echo "Starting deployment..."

# 停止现有服务
echo "Stopping existing service..."
pid=$(lsof -t -i:5173)
if [ ! -z "$pid" ]; then
    kill $pid
fi

# 清理旧的构建文件
echo "Cleaning old build..."
rm -rf dist

# 确保环境变量文件存在
echo "Checking environment files..."
if [ ! -f .env.production ]; then
    echo "Creating .env.production..."
    echo "VITE_API_URL=http://192.168.18.112:8080" > .env.production
fi

# 安装依赖
echo "Installing dependencies..."
npm install

# 打包（修改这行）
echo "Building application..."
npm run build  # 删除了 -- --mode production，因为默认就是 production 模式

# 启动预览服务
echo "Starting preview server..."
nohup npm run preview > server.log 2>&1 &

echo "Deployment completed! Application is running on http://192.168.18.112:5173"