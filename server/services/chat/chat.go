package chat

import (
	"context"
	"log"

	"github.com/bungysheep/chat-app/server/api/chat"
)

type chatServiceServer struct {
	messages chan string
}

// NewChatServiceServer is the Chat service implementation
func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{messages: make(chan string, 100)}
}

func (cs *chatServiceServer) Send(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	if req.GetMessage() != nil {
		if req.GetMessage().GetText() != "" {
			log.Printf("Received message: %s", req.GetMessage().GetText())
			cs.messages <- req.GetMessage().GetText()
		} else {
			log.Printf("Receive message: <empty>")
		}
	} else {
		log.Printf("Received message: nil")
	}
	return nil, nil
}

func (cs *chatServiceServer) Subscribe(req *chat.SubscribeRequest, ss chat.ChatService_SubscribeServer) error {
	for {
		message := <-cs.messages
		if message == "<exit>" {
			return nil
		}

		chatMessage := &chat.Message{
			Text: message,
		}
		if err := ss.Send(&chat.SubscribeResponse{Message: chatMessage}); err != nil {
			cs.messages <- message
			log.Printf("Failed to stream SubscribeResponse: %v", err)
			return err
		}
		log.Printf("Sent message: %v", chatMessage)
	}
}
