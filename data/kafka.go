package data

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func KafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}
	return p
}
