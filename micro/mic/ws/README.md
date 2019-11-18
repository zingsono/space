# Ws消息服务

表名： ms_ws   
字段： 客户端ID、消息ID、消息内容、过期时间、创建时间、状态（0=待处理 1=已处理 9=删除）

ws集群方案： 客户端连接，使用客户端ID作为KEY，消息作为value，通过mongodb watch功能。
## 发送消息


## 订阅消息



## Chanage Log 

- [Feature] 根据服务名数组查询配置集合


### v1.0.0
    [Release 2019-09-11 ]
- [Feature] 创建项目 







