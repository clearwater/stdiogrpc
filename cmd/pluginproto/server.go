package pluginproto

import (
	"context"
	"log"
)

// server is used to implement pluginproto.PluginServer.
type server struct {
	log *log.Logger
	UnimplementedPluginServer
}

func (s *server) CallPlugin(ctx context.Context, in *PluginRequest) (*PluginReply, error) {
	s.log.Println("CallPlugin called ", in)
	return &PluginReply{Message: "I see you " + in.Message}, nil
}

func NewServerImpl(log *log.Logger) *server {
	return &server{log: log}
}
