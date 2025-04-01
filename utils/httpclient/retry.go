package httpclient

import (
	"net/http"
	"time"
)

func RetryRequest(url string, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error
	for attempts := 1; attempts <= maxRetries; attempts++ {
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return resp, nil
		}
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	return nil, err
}
