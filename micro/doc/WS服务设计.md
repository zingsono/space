# WebSocket服务

定义ws协议通信方式

## 请求URL
```
ws:127.0.0.1:5010/ws/gateway?cid=uuid
```
`cid` 参数为客户端编号，用于当前连接唯一标识，日志记录等，登录成功后服务端分配。  
`token` 登录会话，根据token可获取当前用户信息，判断权限

更多客户端数据，用于记录客户端日志信息。

消息类型： REQ=请求  RES=响应 PUSH=推送

HQL_REQ=Graphql请求  
HQL_RES=Graphql响应   
PUSH_ORDER=订单状态推送  
PUSH_MSG=订单状态推送  



## 发送消息
```json
{
  "tid": "uuid，交易编号，消息发送端生成",
  "time": "20191105103500",
  "method": "HQL_REQ",
  "app": "member",
  "msg": "交易报文"
}
```
消息类型: 请求、响应、推送


浏览器端： 发送完消息，记录tid，订阅到消息时，赋值到tid对应的值

## 订阅消息    
```json
{
  "type": "",
  "tid": "uuid，交易编号，消息发送端生成，原样返回",
  "body": "交易报文"
}
```
消息类型： 请求、响应、推送

请求类型的，需要发送效应消息
响应类型的，处理数据即可
推送类型的，处理数据即可


