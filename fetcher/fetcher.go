package fetcher

import "net/http"

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(resp, err)
	}

	return resp, err
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
