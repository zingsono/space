#!/bin/sh

# 需要在运行环境的工作目录执行编译
export GOPROXY=https://goproxy.cn
go build -o docker/mic
echo "Go build successful ^_^ "
