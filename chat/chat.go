package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, hr *HelloRequest) (*HelloResponse, error) {
	log.Printf("Received message body from client: %s", hr.GetName())

	return &HelloResponse{Message: "Hello again " + hr.GetName()}, nil
}
