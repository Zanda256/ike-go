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

type Config struct {
	// Timeout specifies a time limit for requests made by
	Timeout int
	// MaxIdleConns controls the maximum number of idle (keep-alive) connections across all hosts. Zero means no limit.
	MaxIdleConns int
	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle (keep-alive) connections to keep per-host.
	MaxIdleConnsPerHost int
}

func NewClientProvider(cfg Config) *ClientProvider {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			// can configure transport to re-use http connections efficiently
			MaxIdleConns: cfg.MaxIdleConns,
		},
		Timeout: time.Duration(cfg.Timeout) * time.Second,
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
	response.Body.Close() // close body to reuse http connection

	return Response{
		StatusCode: response.StatusCode,
		Headers:    h,
		Body:       resBody,
	}, nil
}
