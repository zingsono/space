# 注册中心

服务地址
`ws://127.0.0.1:5801/ws/discover`  

消息数据结构
```json
{
  "id": "instant",
  "name":"serviceName",
  "host": "127.0.0.0",
  "port": "5811",
  "watch": [""]
}
```


## 注册中心

表名：ms_discover   
字段：服务名、IP、port、过期时间、更新时间、
      name   ip  port  expires  updatedAt
      
客户端心跳：定时心跳更新过期时间，超时由TTL索引剔除服务      
客户端监听：更新事件、删除事件 
     
使用ws通信
     


# Change Log 

## 待实现功能
- [Feature] 


## v1.0.0
    [Release 2019-10-11 ]
- [Feature] 创建项目 

