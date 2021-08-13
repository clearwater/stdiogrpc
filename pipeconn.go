package stdiogrpc

import (
	"io"
)

type PipeConn struct {
	Reader io.ReadCloser
	Writer io.WriteCloser
}

func (p PipeConn) Close() error {
	return p.Writer.Close()
}

func (p PipeConn) Read(buf []byte) (int, error) {
	return p.Reader.Read(buf)
}

func (p PipeConn) Write(buf []byte) (int, error) {
	return p.Writer.Write(buf)
}

func NewPipeConn(reader io.ReadCloser, writer io.WriteCloser) PipeConn {
	return PipeConn{reader, writer}
}
