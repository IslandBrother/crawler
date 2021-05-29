package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "github.com/island-brother/crawler/common"

	"github.com/island-brother/crawler/conn"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Fetch(url string) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	go sendToKafka(KAFKA_TOPIC_FETCHED, &Fetched{URL: url, Content: string(bodyBytes)})
}

func reportError(resp *http.Response, err error) {
	if isBanCase(resp) {
		reportBanError(resp, err)
	} else {
		reportHttpError(resp, err)
	}
}

func isBanCase(resp *http.Response) bool {
	return resp.StatusCode == 429
}

func reportBanError(resp *http.Response, err error) {
	sendToKafka(KAFKA_TOPIC_BANNED, Error{
		URL:        resp.Request.URL.RequestURI(),
		StatusCode: resp.StatusCode,
		Error:      GetErrorString(err),
	})
}

func reportHttpError(resp *http.Response, err error) {
	sendToKafka(KAFKA_TOPIC_HTTP_ERROR, Error{
		URL:        resp.Request.URL.RequestURI(),
		StatusCode: resp.StatusCode,
		Error:      GetErrorString(err),
	})
}

func sendToKafka(topic string, data interface{}) {
	producer := conn.KafkaProducer()

	defer producer.Close()

	kafkaValue, _ := json.Marshal(data)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          kafkaValue,
	}, nil)

	producer.Flush(1 * 1000)
}
