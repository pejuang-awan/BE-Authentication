package service

import (
	"io"
	"net/http"
)

func httpCall(method string, url string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
