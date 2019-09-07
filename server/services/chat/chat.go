package chat

import (
	"context"

	"github.com/bungysheep/chat-app/server/api/chat"
)

type chatServiceServer struct {
}

// NewChatServiceServer is the Chat service implementation
func NewChatServiceServer() chat.ChatServiceServer {
	return &chatServiceServer{}
}

func (*chatServiceServer) Send(ctx context.Context, req *chat.SendRequest) (*chat.SendResponse, error) {
	return nil, nil
}

func (*chatServiceServer) Subscribe(req *chat.SubscribeRequest, ss chat.ChatService_SubscribeServer) error {
	return nil
}
