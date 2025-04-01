package httpclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPostJSON(url string, payload interface{}) (*http.Response, error) {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}
