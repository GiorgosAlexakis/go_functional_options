package main

import (
	"net"
)

type Server struct {
	listener net.Listener
	foo      int
}

func (s *Server) Addr() net.Addr
func (s *Server) Shutdown()
func Foo(value int) func(s *Server) {
	return func(s *Server) {
		s.foo = value
	}
}
func NewServer(addr string, options ...func(*Server)) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := Server{listener: l}
	//go srv.run()
	return &srv, nil
}

func main() {
	_, _ = NewServer("localhost", Foo(1000))
}
