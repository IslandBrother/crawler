package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	_, err := Fetch("http://www.naver.com")

	assert.Equal(t, err, nil)
}
