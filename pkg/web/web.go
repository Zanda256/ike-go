package web

import (
	"bytes"
	"encoding/json"
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
			MaxIdleConnsPerHost: 20, // can configure transport to re-use http connections efficiently
		},
		Timeout: 10 * time.Second,
	}

	return &ClientProvider{client: client}
}

func (cp *ClientProvider) SendRequest(method, endpoint string, reqBody []byte) (Response, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return Response{}, err
	}

	response, err := cp.client.Do(req)
	if err != nil {
		return Response{}, err
	}
	h, err := json.Marshal(response.Header)
	if err != nil {
		return Response{}, err
	}
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: response.StatusCode,
		Headers:    h,
		Body:       resBody,
	}, nil
}
