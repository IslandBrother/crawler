package fetcher

import "net/http"

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		reportError(err)
	}

	return resp, err
}

func reportError(err error) error {
	if isBanCase(err) {
		return nil
	} else {
		return nil
	}
}

func isBanCase(err error) bool {
	return true
}
