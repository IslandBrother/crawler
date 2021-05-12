package fetcher

import "net/http"

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, checkErr(err)
	}

	return resp, nil
}

func checkErr(err error) error {
	return nil
}
