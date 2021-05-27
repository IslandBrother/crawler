package fetcher

import "net/http"
import "github.com/island-brother/crawler/data"
import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
import "io/ioutil"

type Data struct {
	url  string
	html string
}

func Fetch(url string) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	go sendHtml(&Data{url: url, html: string(bodyBytes)})
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

func sendHtml(resp *http.Response) {
	err := sendHtmlToKafka(resp)
	if err != nil {
		sendToParser(resp)
	}
}

func sendHtmlToKafka(resp *http.Response) error {
	topic := "content"
	producer := data.KafkaProducer()

	defer producer.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          bodyBytes,
	}, nil)

	producer.Flush(1 * 1000)

	return nil
}

func sendToParser(resp *http.Response) {
	//grpc will be used
}
