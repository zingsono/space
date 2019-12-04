#!/bin/sh

cd cli

TAG="registry.cn-hangzhou.aliyuncs.com/mszs/gateway:1.0"

docker build -t $TAG  ./
echo "Docker build successful ^_^ "

docker push $TAG
echo "Docker push successful ^_^ "


