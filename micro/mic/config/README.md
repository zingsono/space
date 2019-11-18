# 配置中心（config）

配置中心服务端，提供配置信息统一管理服务。


## 配置中心

表名：ms_config   
字段：服务名、配置JSON内容、备注、更新时间、创建时间

客户端监听：配置更新
使用ws通信

所有服务，通过启动命令参数指定配置中心地址，读取配置。


## 功能说明

配置集合：`db.ms_config`

使用Watch监听配置集合的新增修改删除操作。 Websocket发送更新消息到业务服务。

WebSocket服务：`/ws/config?name=appName`

业务服务发送订阅报文：
```json
["default","mongodb"]
```

配置服务发送配置报文：
```json
{
    "name":{
      "key": {  },
      "keys": {  }
    }
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





