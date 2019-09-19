package chat

import (
	"context"
	"log"

	"github.com/bungysheep/chat-app/server/api/chat"
	"github.com/bungysheep/chat-app/server/models/message"
)

type chatServiceServer struct {
	msg chan *message.Message
}

// NewChatServiceServer is the Chat service implementation
func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{
		msg: make(chan *message.Message, 10),
	}
}

func (cs *chatServiceServer) Send(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	if req.GetMessage() != nil {
		if req.GetMessage().GetText() != "" {
			cs.msg <- message.NewMessage(req.GetMessage(), req.GetMessage().GetSendId(), req.GetSubId())
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

	for {
		msg := <-cs.msg
		if req.GetSubId() == msg.GetSubID() {
			if msg.GetMessage().GetText() == "<exit>" {
				return nil
			}
			if err := ss.Send(&chat.SubscribeResponse{Message: msg.GetMessage()}); err != nil {
				return err
			}
		} else {
			cs.msg <- msg
		}
	}
}
