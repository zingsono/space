#!/bin sh

# 在Golang编译环境中执行此脚本，需要支持golang、docker、git

# 服务参数
NAME="config"
VERSION="1.0"

# 拉取项目源码
git clone https://github.com/zingsono/space.git

# 服务跟路径
HPATH="/space/micro/mic/$NAME"

# 进入服务目录编译Golang服务
cd $HPATH
export GOPROXY=https://goproxy.cn
go build -o ./build/$NAME

# 进入容器构建目录构建Docker容器
cd $HPATH/build
cat > Dockerfile << EOF
FROM mszs/centos:7
COPY ./ /
CMD ["$NAME"]
EOF
docker build -t serve/$NAME:$VERSION ./
docker push  serve/$NAME:$VERSION

# 清理编译输出文件
rm -rf $HPATH/build

echo "Build successful ^_^ "

