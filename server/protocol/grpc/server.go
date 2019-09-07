package grpc

import (
	"context"
	"log"
	"net"

	chatapi "github.com/bungysheep/chat-app/server/api/chat"
	chatservice "github.com/bungysheep/chat-app/server/services/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server is type of gRpc server
type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

// GetGrpcServer returns a gRpc server
func (svr *Server) GetGrpcServer() *grpc.Server {
	return svr.grpcServer
}

// GetListener returns a listener
func (svr *Server) GetListener() net.Listener {
	return svr.listener
}

// RunGrpcServer starts gRpc server
func (svr *Server) RunGrpcServer(ctx context.Context, grpcPort string) error {
	log.Printf("gRpc server is starting...\n")

	lis, err := net.Listen("tcp", "0.0.0.0:"+grpcPort)
	if err != nil {
		return err
	}

	// Define gRpc server options
	opts := []grpc.ServerOption{}

	// Create new gRpc server
	s := grpc.NewServer(opts...)

	// Register Chat service
	chatapi.RegisterChatServiceServer(s, chatservice.NewChatServiceServer())

	// Register gRpc reflection
	reflection.Register(s)

	svr.grpcServer = s
	svr.listener = lis

	log.Printf("gRpc server is listening on port %v...\n", grpcPort)

	return s.Serve(lis)
}
