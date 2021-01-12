package nanotify

import "time"

// CreateWebhookWatcherParams contains the data for createing a webhook watcher
type CreateWebhookWatcherParams struct {
	Account string    `json:"account"`
	URL     string    `json:"url"`
	Event   EventType `json:"event"`
}

// WebhookWatcher encapsulates the data for the webhook watcher
type WebhookWatcher struct {
	ID         string    `json:"id"`
	Account    string    `json:"account"`
	Network    Network   `json:"network"`
	URL        string    `json:"url"`
	Event      EventType `json:"event"`
	InsertedAt time.Time `json:"inserted_at"`
}
