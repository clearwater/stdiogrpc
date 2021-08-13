package helloworld

import (
	"context"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct{ log *log.Logger }

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	s.log.Println("SayHello called ", in)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	s.log.Println("SayHelloAgain called ", in)
	return &HelloReply{Message: "Hello again " + in.Name}, nil
}

func NewServerImpl(log *log.Logger) *server {
	return &server{log: log}
}
