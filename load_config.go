package main

import (
	"encoding/json"
	"log"
	"os"
)

// ServerST struct
type ConfigST struct {
	MQTTclientID string `json:"mqttClientID"`
	MQTTserver   string `json:"mqttServer"`
	MQTTuserName string `json:"mqttUserName"`
	MQTTpassword string `json:"mqttPassword"`
	MQTTtopic    string `json:"mqttTopic"`
}

// 读取配置文件并生成附属结构
func loadConfig(filePath string) *ConfigST {
	var tmp ConfigST
	data, err := os.ReadFile(filePath)
	if err == nil {
		err = json.Unmarshal(data, &tmp)
		if err != nil {
			log.Fatalln(err)
		}

	} else {
		log.Println("读取配置文件失败:", err)
	}
	return &tmp
}
