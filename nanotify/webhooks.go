package nanotify

// WebhookWatcherParams contains the data for createing a webhook watcher
type WebhookWatcherParams struct {
	Account string    `json:"account"`
	URL     string    `json:"url"`
	Event   EventType `json:"event"`
}
