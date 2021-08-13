package hostproto

import (
	"context"
	"log"
)

// server is used to implement pluginproto.PluginServer.
type server struct {
	log *log.Logger
	UnimplementedHostServer
}

func (s *server) CallHost(ctx context.Context, in *HostRequest) (*HostReply, error) {
	s.log.Println("CallHost called ", in)
	return &HostReply{Message: "Hello " + in.Message}, nil
}

func NewServerImpl(log *log.Logger) *server {
	return &server{log: log}
}
