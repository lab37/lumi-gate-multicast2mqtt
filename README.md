# 把绿米网关的亮灯组播转换为mqtt消息 
- 配置文件：config.json
```bash
{
    "mqttClientID": "multicast2mqtt",                   # mqtt的客户端ID
    "mqttServer": "192.168.31.196:1883",                # mqtt服务器地址
    "mqttUserName": "aaaaaaaa",                         # mqtt用户名
    "mqttPassword": "P@ssw0rd",                         # mqtt密码
    "mqttTopic": "homeassistant/security/gate/motion"   # 发生亮灯时发送的mqtt主题
}
```
- 使用方法
```bash
multicast2mqtt  -config  config.json
```
