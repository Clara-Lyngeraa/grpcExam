package main

import (
	"context"
	proto "grpcExam/proto"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedIncrementServiceServer
	name string
	port int
}

var amount int32

func main() {

}

func startServer(server *Server) {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(server.port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	proto.RegisterIncrementServiceServer(grpcServer, server)
	serveError := grpcServer.Serve(listener)
	if serveError != nil {
		log.Fatalf("failed to serve: %v", serveError)
	}
}

func (s *Server) Increment(ctx context.Context, in *proto.IncrementRequest) (*proto.IncrementResponse, error) {
	return &proto.IncrementResponse{Value: amount}, nil
}
