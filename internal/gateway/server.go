package gateway

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/lyounthzzz/realtimex/api/msg"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/operation"
	websocket2 "github.com/lyounthzzz/realtimex/internal/gateway/websocket"
	"github.com/lyounthzzz/realtimex/pkg/middleware"
	"google.golang.org/protobuf/proto"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	wsSrv *websocket2.Server
	msgc  msg.MsgServiceClient

	ms middleware.Middleware
}

func NewServer() *Server {
	srv := &Server{}
	srv.ms = middleware.Chain(middleware.Logger())

	srv.wsSrv = websocket2.NewServer(
		websocket2.AuthnClient(&authClient{}),
	)

	go srv.ListenWebsocketServer(srv.wsSrv)
	return srv
}

func (srv *Server) Start(ctx context.Context) error {
	if err := srv.wsSrv.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (srv *Server) Stop(ctx context.Context) error {
	return srv.wsSrv.Stop(ctx)
}

func (srv *Server) ProtoHandler() middleware.Handler {
	return func(ctx context.Context, proto *protocol.Proto) (proto.Message, error) {
		return srv.ms(srv.handleProto)(ctx, proto)
	}
}

func (srv *Server) ListenWebsocketServer(ws *websocket2.Server) {
	for sess := range ws.Sessions() {
		sess.AddHandler(srv.ProtoHandler())

		go sess.Listen()
	}
}

func (srv *Server) handleProto(ctx context.Context, proto *protocol.Proto) (proto.Message, error) {

	switch proto.Operation {
	case operation.Ping:

	}

	return nil, nil
}
