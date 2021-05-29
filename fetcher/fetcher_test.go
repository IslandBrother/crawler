package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/island-brother/crawler/common"
	"github.com/island-brother/crawler/conn"
	"github.com/stretchr/testify/assert"
)

const TestConsumerGroupId = "fetcher-test"
const TestBasicUrl = "http://www.naver.com"

func TestSendFetchedToKafka(t *testing.T) {
	url := "http://demo.wisetracker.co.kr/test.html"

	fetched := getFetched(url)
	sendToKafka(KAFKA_TOPIC_FETCHED, fetched)

	lastFetched := consumeFetched()

	assert.Equal(t, fetched.URL, lastFetched.URL)
	assert.Equal(t, fetched.Content, lastFetched.Content)
}

func getFetched(url string) *Fetched {
	resp, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fetched := &Fetched{URL: url, Content: string(bodyBytes)}
	return fetched
}

func consumeFetched() Fetched {
	consumer := conn.KafkaConsumer("test", []string{"fetched"})
	msg, err := consumer.ReadMessage(-1)

	if err != nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

	fetched := Fetched{}
	json.Unmarshal(msg.Value, &fetched)
	return fetched
}

func TestFetchWithReturn(t *testing.T) {
	url := TestBasicUrl
	fetched := FetchWithReturn(url)
	fmt.Println(fetched.URL)
}

func TestIsBanCase(t *testing.T) {
	resp, _ := getBanCaseResp(TestBasicUrl)
	result := isBanCase(resp)
	assert.Equal(t, true, result)
}

func TestIsNotBanCase(t *testing.T) {
	resp, _ := getHttpErrorCaseResp(TestBasicUrl)
	result := isBanCase(resp)
	assert.Equal(t, false, result)
}

func TestReportErrorOfBanCase(t *testing.T) {
	topic := "banned"
	resp, err := getBanCaseResp(TestBasicUrl)
	reportError(resp, err)

	checkJustProducedError(t, topic, resp, err)
}

func TestReportErrorOfHttpErrorCase(t *testing.T) {
	topic := "http-error"
	resp, err := getHttpErrorCaseResp(TestBasicUrl)
	reportError(resp, err)

	checkJustProducedError(t, topic, resp, err)
}

func checkJustProducedError(t *testing.T, topic string, resp *http.Response, err error) {
	consumer := conn.KafkaConsumer(TestConsumerGroupId, []string{topic})
	msg, _ := consumer.ReadMessage(-1)
	lastBanned := Error{}
	json.Unmarshal(msg.Value, lastBanned)

	assert.Equal(t, TestBasicUrl, lastBanned.URL)
	assert.Equal(t, resp.StatusCode, lastBanned.StatusCode)
	assert.Equal(t, err.Error(), lastBanned.Error)
}

func getBanCaseResp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	resp.StatusCode = 429
	return resp, err
}

func getHttpErrorCaseResp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	resp.StatusCode = 404
	return resp, err
}
