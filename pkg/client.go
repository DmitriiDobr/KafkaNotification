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

func New(cfg *Config) (*Client, error) {
	conn, err := kafka.Dial("tcp", cfg.Brokers)
	if err != nil {
		return nil, err
	}

	return &Client{connection: conn}, nil
}

func (c *Client) Notify(ctx context.Context, message Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	_, err = c.connection.WriteMessages(kafka.Message{Key: []byte("notifier"), Value: body})

	if err != nil {
		return err
	}
	fmt.Println("Message sent!")
	return nil
}

type Status int

const (
	Success Status = iota
	Warning
	Error
)

type Message struct {
	UserID int    `json:"user_id"`
	Status Status `json:"status"`
	Header string `json:"header"`
	Body   string `json:"body"`
}
