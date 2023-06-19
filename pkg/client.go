package kafkaNotification

import (
	"context"
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

func New(ctx context.Context, cfg *Config) (*Client, error) {
	conn, err := kafka.DialLeader(ctx, "tcp", cfg.Brokers, cfg.Topic, 1)
	if err != nil {
		return nil, err
	}
	return &Client{connection: conn}, nil
}

func (c *Client) Notify(ctx context.Context, message Message, topic string) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = c.connection.WriteMessages(kafka.Message{
		Topic: topic,
		Key:   []byte("notifier"),
		Value: body,
	})

	if err != nil {
		return err
	}
	fmt.Println("Message sent!")
	return nil
}
