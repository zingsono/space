#!/bin sh

export GOPROXY=https://goproxy.cn
go build -o ../../bin/config

cd ../../bin/
mkdir "docker"

# 生成Dockerfile 到docker文件夹
FORM alpine:3.10
COPY . /

# 生成构建镜像脚本到docker文件夹
cat EOF>>>
docker build -t
EOF<<<



