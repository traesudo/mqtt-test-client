package modules

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	Client mqtt.Client
	Config MQTTConfig
}

type MQTTConfig struct {
	Broker   string
	ClientID string
	Username string
	Password string
	Topic    string
	QoS      byte
}

// NewMQTTClient 创建一个新的 MQTT 客户端并连接到服务器
func NewMQTTClient(config MQTTConfig) (*MQTTClient, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Broker)
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MQTTClient{
		Client: client,
		Config: config,
	}, nil
}

func (m *MQTTClient) PublishMessage(message string) error {
	token := m.Client.Publish(m.Config.Topic, m.Config.QoS, false, message)
	token.Wait()
	return token.Error()
}

func (m *MQTTClient) Subscribe(messageHandler mqtt.MessageHandler) error {
	token := m.Client.Subscribe(m.Config.Topic, m.Config.QoS, messageHandler)
	token.Wait()
	return token.Error()
}
