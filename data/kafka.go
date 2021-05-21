package data

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func produce(topic string, data string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()
}
