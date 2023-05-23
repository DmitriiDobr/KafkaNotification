package kafkaNotification

type Message struct {
	UserId int    `json:"user_id"`
	Status string `json:"status"`
	Header string `json:"header"`
	Body   string `json:"body"`
}

const (
	topic          = "my-kafka-topic"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)
