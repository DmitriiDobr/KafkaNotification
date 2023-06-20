package kafkaNotification

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	connection *kafka.Conn
	topic      string
	address    string
}

type Config struct {
	Brokers string
	Topic   string
	Address string
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
	return &Client{connection: conn, topic: cfg.Topic, address: cfg.Address}, nil
}

func (c *Client) Notify(ctx context.Context, message Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP(c.address),
		Topic:    c.topic,
		Balancer: &kafka.LeastBytes{},
	}
	err = w.WriteMessages(ctx, kafka.Message{Key: []byte("notifier"), Value: body})
	if err != nil {
		return err
	}
	fmt.Println("Message sent!")
	return nil
}
