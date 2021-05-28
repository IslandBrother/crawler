package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/island-brother/crawler/data"
	"github.com/stretchr/testify/assert"
)

func TestSendHtmlToKafka(t *testing.T) {

	url := "http://demo.wisetracker.co.kr/test.html"

	resp, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	sendFetchedToKafka(&Fetched{URL: url, Content: string(bodyBytes)})

	consumer := data.KafkaConsumer("test", []string{"fetched"})
	msg, err := consumer.ReadMessage(-1)

	if err != nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

	fetched := Fetched{}
	json.Unmarshal(msg.Value, &fetched)

	assert.Equal(t, url, fetched.URL)
}

func TestReportError(t *testing.T) {
	// _, _ := Fetch("http://invalid")
}

func TestIsBanCase(t *testing.T) {
	resp, _ := http.Get("http://www.naver.com")
	resp.StatusCode = 429
	result := isBanCase(resp)
	assert.Equal(t, true, result)
}

func TestIsNotBanCase(t *testing.T) {
	resp, _ := http.Get("http://www.naver.com")
	result := isBanCase(resp)
	assert.Equal(t, false, result)
}
