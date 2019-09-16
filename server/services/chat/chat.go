package chat

import (
	"context"
	"log"

	"github.com/bungysheep/chat-app/server/api/chat"
	"github.com/bungysheep/chat-app/server/models/subscriber"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type chatServiceServer struct {
	subscriber []*subscriber.Subscriber
}

// NewChatServiceServer is the Chat service implementation
func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{
		subscriber: make([]*subscriber.Subscriber, 10),
	}
}

func (cs *chatServiceServer) Send(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	if req.GetMessage() != nil {
		if req.GetMessage().GetText() != "" {
			sub := cs.subscriber[0]
			if sub != nil && cs.subscriber[0].GetIsActive() {
				log.Printf("Received message: %s", req.GetMessage().GetText())

				if req.GetMessage().GetText() == "<exit>" {
					cs.subscriber[0].SetError(nil)
				}
				if err := cs.subscriber[0].GetSubscriberServer().Send(&chat.SubscribeResponse{Message: req.GetMessage()}); err != nil {
					log.Printf("Failed to stream SubscribeResponse: %v", err)

					cs.subscriber[0].SetError(err)
				}
			} else {
				return nil, status.Errorf(codes.NotFound, "Not found an active chat subscriber")
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

	cs.subscriber[0] = subscriber.NewSubscriber(ss)

	return <-cs.subscriber[0].GetError()
}
