package chat

import (
	"context"
	"log"

	"github.com/bungysheep/chat-app/server/api/chat"
)

type chatServiceServer struct {
	chatSubSvr chat.ChatService_SubscribeServer
	err        chan error
}

// NewChatServiceServer is the Chat service implementation
func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{chatSubSvr: nil, err: make(chan error, 1)}
}

func (cs *chatServiceServer) Send(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	if req.GetMessage() != nil {
		if req.GetMessage().GetText() != "" {
			log.Printf("Received message: %s", req.GetMessage().GetText())

			if req.GetMessage().GetText() == "<exit>" {
				cs.err <- nil
			}
			if err := cs.chatSubSvr.Send(&chat.SubscribeResponse{Message: req.GetMessage()}); err != nil {
				log.Printf("Failed to stream SubscribeResponse: %v", err)

				cs.err <- err
			}
		} else {
			log.Printf("Receive message: <empty>")
		}
	} else {
		log.Printf("Received message: nil")
	}
	return &chat.SendResponse{}, nil
}

func (cs *chatServiceServer) Subscribe(req *chat.SubscribeRequest, ss chat.ChatService_SubscribeServer) error {
	log.Printf("Chat subscription is starting...")

	cs.chatSubSvr = ss

	return <-cs.err
}
