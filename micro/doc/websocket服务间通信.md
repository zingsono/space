
每个服务提供WS服务： `ws://service/ws/graphql`，可选参数sid为会话编号，用于读取当前登录用户。

服务之间通过websocket连接，记录连接与断开日志，实现服务监控。

## 场景：微服务与注册中心的连接

连接地址： `ws://127.0.0.1:8611/ws/discover?name=gateway` 
业务服务发送固定格式消息： 注册信息、订阅服务名 
```
{ name:"discover","addr":"127.0.0.1" }
```


监听Mongodb集合，注册中心主动广播固定格式消息：
```
{ name: ["127.0.0.1:80","127.0.0.1:80"],name2:[],name3:[]}
```
