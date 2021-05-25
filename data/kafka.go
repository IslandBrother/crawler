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

func KafkaConsumer(topics []string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)

	return c
}
