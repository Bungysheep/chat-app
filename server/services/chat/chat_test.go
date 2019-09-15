package chat

import (
	"context"
	"fmt"
	"testing"

	"github.com/bungysheep/chat-app/server/api/chat"
	"github.com/bungysheep/chat-app/server/api/chat/mock"
	"github.com/golang/mock/gomock"
)

func TestSendNilMessage(t *testing.T) {
	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	sendReq := &chat.SendRequest{}
	sendResp, err := chatSvc.Send(ctx, sendReq)
	if err != nil {
		t.Errorf("Expected no Error, but got %v", err)
	}

	if sendResp == nil {
		t.Errorf("Expected a Send Response, but got no Send Response")
	}
}

func TestSendEmptyMessage(t *testing.T) {
	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	sendReq := &chat.SendRequest{Message: &chat.Message{Text: ""}}
	sendResp, err := chatSvc.Send(ctx, sendReq)
	if err != nil {
		t.Errorf("Expected no Error, but got %v", err)
	}

	if sendResp == nil {
		t.Errorf("Expected a Send Response, but got no Send Response")
	}
}

func TestSendAMessage(t *testing.T) {
	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	sendReq := &chat.SendRequest{Message: &chat.Message{Text: "Hello World!"}}
	sendResp, err := chatSvc.Send(ctx, sendReq)
	if err != nil {
		t.Errorf("Expected no Error, but got %v", err)
	}

	if sendResp == nil {
		t.Errorf("Expected a Send Response, but got no Send Response")
	}
}

func TestSendAndSubscribeMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

	go func() {
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{Text: "Hello World!"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{Text: "This is gRpc"}}).Return(nil)

		chatSvc.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "Hello World!"}})
		chatSvc.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "This is gRpc"}})

		chatSvc.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "<exit>"}})
	}()

	if err := chatSvc.Subscribe(&chat.SubscribeRequest{}, mChatSubSvc); err != nil {
		t.Errorf("Expected no Error, but got %v", err)
	}
}

func TestSendAndSubscribeMessageWhenStreamResponseFail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

	go func() {
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{Text: "Hello World!"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{Text: "This is gRpc"}}).Return(fmt.Errorf("Unable to stream message 'This is gRpc'"))

		chatSvc.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "Hello World!"}})
		chatSvc.Send(ctx, &chat.SendRequest{Message: &chat.Message{Text: "This is gRpc"}})
	}()

	if err := chatSvc.Subscribe(&chat.SubscribeRequest{}, mChatSubSvc); err == nil {
		t.Errorf("Expected a Error, but got no Error")
	}
}
