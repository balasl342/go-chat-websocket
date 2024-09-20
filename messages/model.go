package messages

import "time"

// Message represents a chat message
type Message struct {
	Username string    `json:"username"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}
