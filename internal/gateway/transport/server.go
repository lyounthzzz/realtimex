package transport

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/middleware"
)

type (
	Pusher interface {
		Push(ctx context.Context, proto *protocol.Proto) error
	}
	Receiver interface {
		Receive(ctx context.Context, handler middleware.Handler) error
	}
)

type Conn interface {
	transport.Server
	Pusher
	Receiver
	AddMiddleware(ms ...middleware.Middleware)
	GetUser() *gateway.User
}

type Server interface {
	transport.Server
	Connections() chan Conn
}

type connKey struct{}

func ContextWithConn(ctx context.Context, conn Conn) context.Context {
	return context.WithValue(ctx, connKey{}, conn)
}

func ConnFromContext(ctx context.Context) Conn {
	v := ctx.Value(connKey{})
	if v == nil {
		return nil
	}
	return v.(Conn)
}
