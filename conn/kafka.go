package conn

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const bootstrapServers = "localhost"

func KafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		panic(err)
	}
	return p
}

func KafkaConsumer(groupId string, topics []string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics(topics, nil)

	return c
}
