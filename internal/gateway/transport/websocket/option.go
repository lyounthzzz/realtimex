package websocket

import "github.com/gorilla/websocket"

type Option func(*Server)

func Address(addr string) Option {
	return func(srv *Server) { srv.address = addr }
}

func Path(path string) Option {
	return func(srv *Server) { srv.path = path }
}

func Upgrader(upgrader *websocket.Upgrader) Option {
	return func(srv *Server) { srv.upgrader = upgrader }
}
