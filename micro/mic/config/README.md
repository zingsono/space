# 配置中心（config）

配置中心服务端，提供配置信息统一管理服务。


## 功能说明

配置集合：`db.ms_config`

使用Watch监听配置集合的新增修改删除操作。 Websocket发送更新消息到业务服务。

WebSocket服务：`/ws/config`

业务服务发送订阅报文：
```json
{
  "serviceName": ["default","mongodb","serviceName"]
}
```

配置服务发送配置报文：
```json
{
  "key": {  },
  "keys": {  }
}
```


## 项目目录

- doc 项目文档
- mic 服务端数据接口
- toa 运营管理界面



Change Log 

## 待实现功能
- [Feature] 根据服务名数组查询配置集合


## v1.0.0
    [Release 2019-09-11 ]
- [Feature] 创建项目 





