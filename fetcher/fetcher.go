package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/island-brother/crawler/data"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Fetched struct {
	URL     string
	Content string
}

func Fetch(url string) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	go sendFetched(&Fetched{URL: url, Content: string(bodyBytes)})
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

func sendFetched(fetched *Fetched) {
	err := sendFetchedToKafka(fetched)
	if err != nil {
		sendFetchedToParser(fetched)
	}
}

func sendFetchedToKafka(fetched *Fetched) error {
	topic := "fetched"
	producer := data.KafkaProducer()

	defer producer.Close()

	value, _ := json.Marshal(fetched)
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	producer.Flush(1 * 1000)

	return nil
}

func sendFetchedToParser(fetched *Fetched) {
	//grpc will be used
}
