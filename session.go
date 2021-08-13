package stdiogrpc

import (
	"context"
	"net"
	"os"
	"os/exec"

	"github.com/hashicorp/yamux"
)

type Session struct {
	PipeConn PipeConn
	*yamux.Session
}

func (s Session) Dial(context.Context, string) (net.Conn, error) {
	return s.Session.Open()
}

func NewHostSession(cmd *exec.Cmd) (*Session, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	conn := NewPipeConn(stdout, stdin)
	session, err := yamux.Client(conn, nil)
	return &Session{PipeConn: conn, Session: session}, err
}

func NewPluginSession() (*Session, error) {
	conn := NewPipeConn(os.Stdin, os.Stdout)
	session, err := yamux.Server(conn, nil)
	return &Session{PipeConn: conn, Session: session}, err
}
