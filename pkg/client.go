package kafkaNotification

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	connection *kafka.Conn
}

type Config struct {
	Brokers string
	Topic   string
}

func New(cfg *Config) (*Client, error) {
	conn, err := kafka.Dial("tcp", cfg.Brokers)
	if err != nil {
		return nil, err
	}
	topicConfigs := kafka.TopicConfig{Topic: cfg.Topic, NumPartitions: 1, ReplicationFactor: 1}
	err = conn.CreateTopics(topicConfigs)
	if err != nil {
		return nil, err
	}
	return &Client{connection: conn}, nil
}

func (c *Client) Notify(message Message, topic string) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = c.connection.WriteMessages(kafka.Message{
		Key:   []byte("notifier"),
		Value: body,
	})

	if err != nil {
		return err
	}
	fmt.Println("Message sent!")
	return nil
}
