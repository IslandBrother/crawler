package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	_, err := Fetch("http://www.naver.com")
	assert.Equal(t, err, nil)
}

func TestReportError(t *testing.T) {
	// _, _ := Fetch("http://invalid")
}

func TestIsBanCase(t *testing.T) {
	resp, _ := Fetch("http://www.naver.com")
	resp.StatusCode = 429
	result := isBanCase(resp)
	assert.Equal(t, true, result)
}

func TestIsNotBanCase(t *testing.T) {
	resp, _ := Fetch("http://www.naver.com")
	result := isBanCase(resp)
	assert.Equal(t, false, result)
}
