package nanotify

// EventType represents an event enum
type EventType string

const (
	// EventPendingReceive describes the event where a send block
	// is generated with the destination of the send being for the account.
	// This is typically the event you would want if you are wanting
	// to be notified of an incoming transaction to the watched account.
	EventPendingReceive EventType = "pending_receive"

	// EventReceiveBlock describes the event where a receive block
	// is generated for an account.
	EventReceiveBlock EventType = "receive_block"

	// EventSendBlock describes the event where a send block is generated
	// for an account.
	EventSendBlock EventType = "send_block"
)
