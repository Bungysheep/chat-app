package grpc

import (
	"context"
	"log"
	"net"

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

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	reflection.Register(s)

	svr.grpcServer = s
	svr.listener = lis

	log.Printf("gRpc server is listening on port %v...\n", grpcPort)

	return s.Serve(lis)
}
