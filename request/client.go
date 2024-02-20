package request

import (
	"net/http"
	"net/url"
)

const API_PATH = "https://gateway.zeabur.com/graphql"

type Client struct {
	apiKey     string
	apiPath    *url.URL
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	url, err := url.Parse(API_PATH)
	if err != nil {
		panic(err)
	}
	return &Client{
		apiKey:     apiKey,
		apiPath:    url,
		httpClient: &http.Client{},
	}
}