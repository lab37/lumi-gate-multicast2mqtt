package main

import (
	"context"
	"log"
	"time"

	"github.com/lucacasonato/mqtt"
)

func newMqttClient(serverAddr string, username string, password string, clientID string) *mqtt.Client {
	mqttClient, err := mqtt.NewClient(mqtt.ClientOptions{
		// 必须项
		Servers: []string{
			serverAddr,
		},

		// 可选项
		ClientID:      clientID,
		Username:      username,
		Password:      password,
		AutoReconnect: true,
	})
	if err != nil {
		log.Println("can not connect with mqttServer:", err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	//如果在1s内建立了连接则返回nil, 不然就是超时错误了.
	err = mqttClient.Connect(ctx)
	if err != nil {
		log.Println("can not connect with mqttServer:", err)
		panic(err)
	}
	return mqttClient
}

func mqttPubWithTimeout(mqttClient *mqtt.Client, topic string, payload string, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	err := mqttClient.PublishString(ctx, topic, payload, mqtt.AtLeastOnce)
	if err != nil {
		log.Println("failed to publish to config service:", err)
	}
}

func mqttSubWithTimeout(mqttClient *mqtt.Client, topic string, duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	err := mqttClient.Subscribe(ctx, topic, mqtt.AtMostOnce)
	if err != nil {
		log.Println("failed to subscribe to config service:", err)
	}
}
