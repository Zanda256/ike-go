package web

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type ClientProvider struct {
	client *http.Client
}

func NewClientProvider() *ClientProvider {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20, // can be read from config
		},
		Timeout: 10 * time.Second, // can be read from config
	}

	return &ClientProvider{client: client}
}

func (cp *ClientProvider) SendRequest(method, endpoint string, reqBody []byte) ([]byte, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	response, err := cp.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
