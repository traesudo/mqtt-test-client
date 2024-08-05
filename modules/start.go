package modules

import (
	"errors"
	"fmt"
	"log"
)

type Config struct {
	MQTTConfig       MQTTConfig
	ConcurrencyLimit int
	Message          string
}

func (c *Config) Verify() error {
	log.Println("check all params password ", c.MQTTConfig.Password)
	log.Println("check all params username ", c.MQTTConfig.Username)
	log.Println("check all params clientID ", c.MQTTConfig.ClientID)
	log.Println("check all params broker ", c.MQTTConfig.Broker)
	log.Println("check all params topic ", c.MQTTConfig.Topic)

	if len(c.MQTTConfig.Password) == 0 || len(c.MQTTConfig.Username) == 0 || len(c.MQTTConfig.ClientID) == 0 || len(c.MQTTConfig.Broker) == 0 || len(c.MQTTConfig.Topic) == 0 {
		return errors.New("参数错误")
	}
	if c.ConcurrencyLimit < 0 {
		return errors.New("并发参数错误")
	}
	if len(c.Message) == 0 {
		return errors.New("消息不能为空")
	}
	return nil
}

func (c *Config) Stop() bool {
	return StopTaskActuator()
}
func (c *Config) Start() bool {

	err := c.Verify()
	if err != nil {
		log.Println("verify err:", err)
		return false
	}

	mqttClient, err := NewMQTTClient(c.MQTTConfig)
	if err != nil {
		log.Println("init client err:", err)
	}
	err = mqttClient.PublishMessage(c.Message)
	if err != nil {
		fmt.Println("err:", err)
		return false
	}
	CurrentConfig.MQTTConfig = c.MQTTConfig
	CurrentConfig.Message = c.Message
	CurrentConfig.ConcurrencyLimit = c.ConcurrencyLimit
	fmt.Println("查看 CurrentConfig", CurrentConfig)
	StartTaskActuator()
	return true
}
