package main

import (
	"encoding/json"
	"flag"
	"time"
)

func main() {

	var filePath string
	flag.StringVar(&filePath, "config", `/root/user-systemd-services/multicast2mqtt/mqtt_client.json`, "配置文件路径")
	flag.Parse()

	// Config global
	var Config = loadConfig(filePath)
	// 初始化mqtt客户端
	mqttClient := newMqttClient(Config.MQTTserver, Config.MQTTuserName, Config.MQTTpassword, Config.MQTTclientID)

	// 下面这段只为了实时性这里采用监听绿米网关组播消息的方式来触发动作
	multicastMessageCh := make(chan message, 20)
	go udpMulticastReceiver("224.0.0.50", 9898, "", multicastMessageCh)
	go func() {
		for message := range multicastMessageCh {

			multicastData := dataST{}
			multicastPayload := payload{}
			json.Unmarshal(message.Data, &multicastData)
			json.Unmarshal([]byte(multicastData.Payload), &multicastPayload)
			if multicastPayload.RGB > 0 {
				// 推流命令执行以后发送MQTT消息通知
				mqttPubWithTimeout(mqttClient, Config.MQTTtopic, "ok", 1*time.Second)
			}

		}
	}()
	// mqttSubWithTimeout(mqttClient, "homeassistant/security/gate/motion", 1*time.Second)
	//  卡住主协程不要退出
	select {}
}
