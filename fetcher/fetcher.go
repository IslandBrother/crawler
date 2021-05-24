package fetcher

import "net/http"
import "github.com/island-brother/crawler/data"
import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

func Fetch(url string) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	sendHtml(resp)
}

func sendHtml(resp *http.Response){
	err := sendHtmlToKafka(resp)
	if(err != nil){
		sendToParserDirectly(resp)
	}
}

func sendHtmlToKafka(resp *http.Response){
	topic := "html"
	producer := data.KafkaProducer();

	defer producer.Close()
	
	bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	producer.Produce(&kafka.Message{
		TopicPartition: kafaka.TopicPartition{Topic: &fetched, Partition: kafka.PartitionAny},
		value:string(bodyBytes)
	})

	producer.Flush(15 * 1000)
}

func sendToParserDirectly(resp *http.Response){
	//grpc will be used
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
