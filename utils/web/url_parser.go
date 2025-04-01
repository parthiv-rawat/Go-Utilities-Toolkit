package web

import (
	"fmt"
	"net/http"
)

func CheckURLStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "UNREACHABLE"
	}
	defer resp.Body.Close()
	return fmt.Sprintf("Status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
}
