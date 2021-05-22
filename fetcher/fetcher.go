package fetcher

import "net/http"
import "github.com/island-brother/crawler/data"
import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	topic := "fetched"
	producer := data.KafkaProducer();
	producer.Produce(&kafka.Message{
		TopicPartition: kafaka.TopicPartition{Topic: &fetched, Partition: kafka.PartitionAny},
		value:[]byte(resp)
	})

	return resp, err
}

func reportError(resp *http.Response, err error) {
	if isBanCase(resp) {
		reportBanError(err)
	} else {
		reportHttpError(err)
	}
}

func isBanCase(resp *http.Response) bool {
	return resp.StatusCode == 429
}

func reportBanError(err error) {

}

func reportHttpError(err error) {

}
