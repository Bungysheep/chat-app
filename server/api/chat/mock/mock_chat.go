// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/chat/chat.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	chat "github.com/bungysheep/chat-app/server/api/chat"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockChatServiceClient is a mock of ChatServiceClient interface
type MockChatServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceClientMockRecorder
}

// MockChatServiceClientMockRecorder is the mock recorder for MockChatServiceClient
type MockChatServiceClientMockRecorder struct {
	mock *MockChatServiceClient
}

// NewMockChatServiceClient creates a new mock instance
func NewMockChatServiceClient(ctrl *gomock.Controller) *MockChatServiceClient {
	mock := &MockChatServiceClient{ctrl: ctrl}
	mock.recorder = &MockChatServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatServiceClient) EXPECT() *MockChatServiceClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockChatServiceClient) Send(ctx context.Context, in *chat.SendRequest, opts ...grpc.CallOption) (*chat.SendResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*chat.SendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockChatServiceClientMockRecorder) Send(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatServiceClient)(nil).Send), varargs...)
}

// Subscribe mocks base method
func (m *MockChatServiceClient) Subscribe(ctx context.Context, in *chat.SubscribeRequest, opts ...grpc.CallOption) (chat.ChatService_SubscribeClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(chat.ChatService_SubscribeClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockChatServiceClientMockRecorder) Subscribe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockChatServiceClient)(nil).Subscribe), varargs...)
}

// MockChatService_SubscribeClient is a mock of ChatService_SubscribeClient interface
type MockChatService_SubscribeClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_SubscribeClientMockRecorder
}

// MockChatService_SubscribeClientMockRecorder is the mock recorder for MockChatService_SubscribeClient
type MockChatService_SubscribeClientMockRecorder struct {
	mock *MockChatService_SubscribeClient
}

// NewMockChatService_SubscribeClient creates a new mock instance
func NewMockChatService_SubscribeClient(ctrl *gomock.Controller) *MockChatService_SubscribeClient {
	mock := &MockChatService_SubscribeClient{ctrl: ctrl}
	mock.recorder = &MockChatService_SubscribeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatService_SubscribeClient) EXPECT() *MockChatService_SubscribeClientMockRecorder {
	return m.recorder
}

// Recv mocks base method
func (m *MockChatService_SubscribeClient) Recv() (*chat.SubscribeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*chat.SubscribeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockChatService_SubscribeClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).Recv))
}

// Header mocks base method
func (m *MockChatService_SubscribeClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockChatService_SubscribeClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).Header))
}

// Trailer mocks base method
func (m *MockChatService_SubscribeClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockChatService_SubscribeClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).Trailer))
}

// CloseSend mocks base method
func (m *MockChatService_SubscribeClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockChatService_SubscribeClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockChatService_SubscribeClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChatService_SubscribeClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChatService_SubscribeClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChatService_SubscribeClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChatService_SubscribeClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChatService_SubscribeClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_SubscribeClient)(nil).RecvMsg), m)
}

// MockChatServiceServer is a mock of ChatServiceServer interface
type MockChatServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatServiceServerMockRecorder
}

// MockChatServiceServerMockRecorder is the mock recorder for MockChatServiceServer
type MockChatServiceServerMockRecorder struct {
	mock *MockChatServiceServer
}

// NewMockChatServiceServer creates a new mock instance
func NewMockChatServiceServer(ctrl *gomock.Controller) *MockChatServiceServer {
	mock := &MockChatServiceServer{ctrl: ctrl}
	mock.recorder = &MockChatServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatServiceServer) EXPECT() *MockChatServiceServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockChatServiceServer) Send(arg0 context.Context, arg1 *chat.SendRequest) (*chat.SendResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*chat.SendResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockChatServiceServerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatServiceServer)(nil).Send), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockChatServiceServer) Subscribe(arg0 *chat.SubscribeRequest, arg1 chat.ChatService_SubscribeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockChatServiceServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockChatServiceServer)(nil).Subscribe), arg0, arg1)
}

// MockChatService_SubscribeServer is a mock of ChatService_SubscribeServer interface
type MockChatService_SubscribeServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatService_SubscribeServerMockRecorder
}

// MockChatService_SubscribeServerMockRecorder is the mock recorder for MockChatService_SubscribeServer
type MockChatService_SubscribeServerMockRecorder struct {
	mock *MockChatService_SubscribeServer
}

// NewMockChatService_SubscribeServer creates a new mock instance
func NewMockChatService_SubscribeServer(ctrl *gomock.Controller) *MockChatService_SubscribeServer {
	mock := &MockChatService_SubscribeServer{ctrl: ctrl}
	mock.recorder = &MockChatService_SubscribeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChatService_SubscribeServer) EXPECT() *MockChatService_SubscribeServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockChatService_SubscribeServer) Send(arg0 *chat.SubscribeResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockChatService_SubscribeServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).Send), arg0)
}

// SetHeader mocks base method
func (m *MockChatService_SubscribeServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockChatService_SubscribeServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).SetHeader), arg0)
}

// SendHeader mocks base method
func (m *MockChatService_SubscribeServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockChatService_SubscribeServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).SendHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockChatService_SubscribeServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockChatService_SubscribeServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).SetTrailer), arg0)
}

// Context mocks base method
func (m *MockChatService_SubscribeServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockChatService_SubscribeServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).Context))
}

// SendMsg mocks base method
func (m_2 *MockChatService_SubscribeServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockChatService_SubscribeServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).SendMsg), m)
}

// RecvMsg mocks base method
func (m_2 *MockChatService_SubscribeServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockChatService_SubscribeServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockChatService_SubscribeServer)(nil).RecvMsg), m)
}
