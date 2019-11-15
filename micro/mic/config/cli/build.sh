#!/bin sh

# 服务参数
NAME="config"
VERSION="1.0"

# 编译Golang服务
cd ../
export GOPROXY=https://goproxy.cn
go build -o ./build/$NAME

# 创建Dockerfile
cat > Dockerfile << EOF
FROM mszs/centos:7
COPY ./ /
CMD ["$NAME"]
EOF

# 构建Docker镜像推送到镜像仓库
cd ./build
docker build -t serve/$NAME:$VERSION
docker push  serve/$NAME:$VERSION





