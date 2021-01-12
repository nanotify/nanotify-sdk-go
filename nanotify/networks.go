package nanotify

// Network is the Node network that is running
type Network string

const (
	// NetworkNano is the live Nano network
	NetworkNano Network = "nano"

	// NetworkNanoBeta is the beta Nano network
	NetworkNanoBeta Network = "nano-beta"
)
