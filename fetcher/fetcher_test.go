package fetcher

import (
	"fmt"
	"github.com/island-brother/crawler/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestSendHtmlToKafka(t *testing.T) {
	resp, _ := http.Get("http://www.naver.com")
	sendHtmlToKafka(resp)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	consumer := data.KafkaConsumer("test", []string{"content"})
	msg, err := consumer.ReadMessage(-1)

	if err != nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

	assert.Equal(t, bodyBytes, msg.Value)
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
