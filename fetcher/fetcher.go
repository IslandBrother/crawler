package fetcher

import "net/http"

func Fetch(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		go reportError(err)
	}

	return resp, err
}

func reportError(err error) {
	if isBanCase(err) {
		reportBanError(err)
	} else {
		reportHttpError(err)
	}
}

func isBanCase(err error) bool {
	return true
}

func reportBanError(err error) {

}

func reportHttpError(err error) {

}
