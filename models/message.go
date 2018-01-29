package models

// MessageState defines the state of a message
type MessageState int

// Enumeration for MessageState
const (
	OPEN MessageState = 1 + iota
	ACKED
	CLOSED
)

// Message defines a message sent by end user
type Message struct {
	ID        string       `json:"id"`   // message id
	Body      string       `json:"body"` // message body
	State     MessageState // message state
	ServiceID string       `json:"service_id"` // message's service id
	UserID    string       `json:"user_id"`    // message's user id
}
