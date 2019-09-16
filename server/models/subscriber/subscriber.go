package subscriber

import "github.com/bungysheep/chat-app/server/api/chat"

// Subscriber is type of Chat Subscriber
type Subscriber struct {
	chatSubSvr chat.ChatService_SubscribeServer
	err        chan error
	isActive   bool
}

// NewSubscriber creates an instance of Chat Subscriber
func NewSubscriber(subSvr chat.ChatService_SubscribeServer) *Subscriber {
	return &Subscriber{
		chatSubSvr: subSvr,
		err:        make(chan error),
		isActive:   true,
	}
}

// GetSubscriberServer returns Chat Subscribe server
func (sub *Subscriber) GetSubscriberServer() chat.ChatService_SubscribeServer {
	return sub.chatSubSvr
}

// SetSubscriberServer sets Chat Subscribe server
func (sub *Subscriber) SetSubscriberServer(subSvr chat.ChatService_SubscribeServer) {
	sub.chatSubSvr = subSvr
}

// GetError returns error
func (sub *Subscriber) GetError() <-chan error {
	return sub.err
}

// SetError sets error
func (sub *Subscriber) SetError(err error) {
	sub.err <- err
}

// GetIsActive returns status of Chat Subscriber
func (sub *Subscriber) GetIsActive() bool {
	return sub.isActive
}

// SetIsActive sets status of Chat Subscriber
func (sub *Subscriber) SetIsActive(isActive bool) {
	sub.isActive = isActive
}
