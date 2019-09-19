package chat

import (
	"context"
	"fmt"
	"testing"
	"time"

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

func TestSendAndSubscribeMessage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

	done := make(chan bool)
	go func() {
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "Hello World!"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "This is gRpc"}}).Return(nil)

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "001"}, mChatSubSvc); err != nil {
			t.Errorf("Expected no Error, but got %v", err)
		}
		done <- true
	}()

	time.Sleep(1 * time.Second)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "001", Text: "Hello World!"}})
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "001", Text: "This is gRpc"}})
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{Text: "<exit>"}})

	<-done
}

func TestSendAndSubscribeMessageWhenStreamResponseFail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

	done := make(chan bool)
	go func() {
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "Hello World!"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "This is gRpc"}}).Return(fmt.Errorf("Unable to stream message 'This is gRpc'"))

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "001"}, mChatSubSvc); err == nil {
			t.Errorf("Expected an Error, but got no Error")
		}
		done <- true
	}()

	time.Sleep(1 * time.Second)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "001", Text: "Hello World!"}})
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "001", Text: "This is gRpc"}})

	<-done
}

func TestSendMessageToOtherSubscriber(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	done := make(chan bool, 2)
	go func() {
		mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "002", Text: "This is gRpc"}}).Return(nil)

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "001"}, mChatSubSvc); err != nil {
			t.Errorf("Expected no Error, but got %v", err)
		}
		done <- true
	}()

	go func() {
		mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "Hello World!"}}).Return(nil)

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "002"}, mChatSubSvc); err != nil {
			t.Errorf("Expected no Error, but got %v", err)
		}
		done <- true
	}()

	time.Sleep(1 * time.Second)

	chatSvc.Send(ctx, &chat.SendRequest{SubId: "002", Message: &chat.Message{SendId: "001", Text: "Hello World!"}})
	time.Sleep(30 * time.Millisecond)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "002", Text: "This is gRpc"}})

	chatSvc.Send(ctx, &chat.SendRequest{SubId: "002", Message: &chat.Message{Text: "<exit>"}})
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{Text: "<exit>"}})

	<-done
	<-done
}

func TestSendMessageToOtherSubscriberWhenStreamResponseFail(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()

	chatSvc := NewChatServiceServer()

	done := make(chan bool, 2)
	go func() {
		mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "002", Text: "Hello World! 002"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "002", Text: "This is gRpc 002"}}).Return(fmt.Errorf("Unable to stream message 'This is gRpc'"))

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "001"}, mChatSubSvc); err == nil {
			t.Errorf("Expected no Error, but got %v", err)
		}
		done <- true
	}()

	go func() {
		mChatSubSvc := mock.NewMockChatService_SubscribeServer(ctl)

		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "Hello World! 001"}}).Return(nil)
		mChatSubSvc.EXPECT().Send(&chat.SubscribeResponse{Message: &chat.Message{SendId: "001", Text: "This is gRpc 001"}}).Return(fmt.Errorf("Unable to stream message 'This is gRpc'"))

		if err := chatSvc.Subscribe(&chat.SubscribeRequest{SubId: "002"}, mChatSubSvc); err == nil {
			t.Errorf("Expected no Error, but got %v", err)
		}
		done <- true
	}()

	time.Sleep(1 * time.Second)

	chatSvc.Send(ctx, &chat.SendRequest{SubId: "002", Message: &chat.Message{SendId: "001", Text: "Hello World! 001"}})
	time.Sleep(30 * time.Millisecond)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "002", Text: "Hello World! 002"}})
	time.Sleep(30 * time.Millisecond)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "002", Message: &chat.Message{SendId: "001", Text: "This is gRpc 001"}})
	time.Sleep(30 * time.Millisecond)
	chatSvc.Send(ctx, &chat.SendRequest{SubId: "001", Message: &chat.Message{SendId: "002", Text: "This is gRpc 002"}})

	<-done
	<-done
}
