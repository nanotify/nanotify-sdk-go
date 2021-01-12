// Package nanotify contains the main nanotify client code
package nanotify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Client is the main point of entry for Nanotify
type Client struct {
	APIKey  string
	Network Network
}

// NewClient will create a new client which authenticates using the associated
// apiKey.
func NewClient(apiKey string, network Network) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("invalid api key")
	}

	if network != NetworkNano && network != NetworkNanoBeta {
		return nil, fmt.Errorf("invalid network")
	}

	return &Client{
		APIKey:  apiKey,
		Network: network,
	}, nil
}

func nanotifyURL() string {
	base := os.Getenv("NANOTIFY_BASE_URL")
	if base == "" {
		return "https://api.nanotify.io"
	}

	return base
}

func urlWithPath(path string) string {
	return nanotifyURL() + "/v1" + path
}

func (c *Client) authenticateRequest(r *http.Request) {
	r.Header.Add("x-api-key", c.APIKey)
}

func (c *Client) errorFromResp(r *http.Response) error {
	if r.StatusCode >= 400 {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", body)
	}
	return nil
}

// CreateWebhookWatcher registers a new Webhook Watcher with the Nanotify
// service.
func (c *Client) CreateWebhookWatcher(
	ctx context.Context, params CreateWebhookWatcherParams,
) (WebhookWatcher, error) {
	result := WebhookWatcher{}
	url := urlWithPath("/webhooks")

	body, err := json.Marshal(params)
	if err != nil {
		return result, fmt.Errorf("failed to marshal json: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, url, bytes.NewReader(body),
	)
	if err != nil {
		return result, fmt.Errorf("failed to create a request: %w", err)
	}

	c.authenticateRequest(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, fmt.Errorf("failed to perform request: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	err = c.errorFromResp(res)
	if err != nil {
		return result, fmt.Errorf("error in response: %w", err)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("error decoding response body: %w", err)
	}

	return result, nil
}
