package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func Producer(ctx context.Context, status status, userId uuid, header, body string) {

	w := kafka.Writer{
		Addr:  kafka.TCP(broker1Address, broker2Address, broker3Address),
		Topic: topic,
	}

	msg := Message{
		UserId: userId,
		Status: status,
		Header: header,
		Body:   body,
	}

	bytesMsg, err := json.Marshal(msg)
	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte("MessageStatus " + string(rune(userId))),
		Value: bytesMsg,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent!")

}
