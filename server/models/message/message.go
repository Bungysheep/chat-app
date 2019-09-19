package message

import "github.com/bungysheep/chat-app/server/api/chat"

// Message is type of Chat Message
type Message struct {
	message *chat.Message
	sendID  string
	subID   string
}

// NewMessage creates an instance of Chat Message
func NewMessage(message *chat.Message, sendID string, subID string) *Message {
	return &Message{
		message: message,
		sendID:  sendID,
		subID:   subID,
	}
}

// GetMessage returns Chat Message
func (sub *Message) GetMessage() *chat.Message {
	return sub.message
}

// GetSendID returns Send ID
func (sub *Message) GetSendID() string {
	return sub.sendID
}

// GetSubID returns Subscriber ID
func (sub *Message) GetSubID() string {
	return sub.subID
}
