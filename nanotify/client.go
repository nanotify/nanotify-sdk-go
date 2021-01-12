// Package nanotify contains the main nanotify client code
package nanotify

import (
	"fmt"
	"net/http"
)

// Client is the main point of entry for Nanotify
type Client struct {
	APIKey     string
	Network    Network
	httpClient *http.Client
}

// NewClient will create a new client which authenticates using the associated
// apiKey
func NewClient(apiKey string, network Network) (*Client, error) {

	if apiKey == "" {
		return nil, fmt.Errorf("invalid api key")
	}

	if network != NetworkNano && network != NetworkNanoBeta {
		return nil, fmt.Errorf("invalid network")
	}

	return &Client{
		APIKey:     apiKey,
		Network:    network,
		httpClient: &http.Client{},
	}, nil
}

// CreateWebhookWatcher registers a new Webhook Watcher with the Nanotify
// service.
func (c *Client) CreateWebhookWatcher(params WebhookWatcherParams) error {
	return nil
}
