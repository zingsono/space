# 微服务网关

所有服务统一入口。

请求报文：
```text
query{
    member{
    
    },
}
```


响应报文：



##部署说明

编译镜像：docker pull golang:1.13-alpine
运行环境：docker pull mszs/alpine:3.10

运行容器：
```
docker run --name ms.gateway --network cluster -p 39701:7908 -dit registry.cn-hangzhou.aliyuncs.com/mszs/gateway:1.0
```

测试环境接口地址：http://211.152.57.29:39701/api/v2/graphql  

## Chanage Log 

待实现功能
- [Feature] 根据服


### v1.0.0 
    [Release 2019-09-11 ]
- [Feature] 创建项目 















































































































