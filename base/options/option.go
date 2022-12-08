package options

import (
	"errors"
	"time"
)

// option 模式
type oClient struct {
	url           string
	optionTimeout time.Duration
}

type options = func(client *oClient)

func NewClient(opts ...options) (*oClient, error) {
	newObj := &oClient{
		url:           "",
		optionTimeout: 30 * time.Second,
	}

	for _, opt := range opts {
		opt(newObj)
	}

	if newObj.url == "" {
		return nil, errors.New("URL is empty")
	}

	return newObj, nil
}

func WithURL(url string) options {
	return func(client *oClient) {
		client.url = url
	}
}

func WithOptionTimeout(timeout time.Duration) options {
	return func(client *oClient) {
		client.optionTimeout = timeout
	}
}
