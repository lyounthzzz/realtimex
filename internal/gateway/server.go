package gateway

import (
	"context"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport/mqtt"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport/websocket"
	"github.com/lyounthzzz/realtimex/pkg/middleware"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	wsSrv   *websocket.Server
	mqttSrv *mqtt.Server
	ms      middleware.Middleware
}

func (srv *Server) AddWebsocketServer() {

}

func (srv *Server) AddMqttServer() {

}

func (srv *Server) Start(ctx context.Context) error {
	if srv.wsSrv != nil {
		if err := srv.wsSrv.Start(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (srv *Server) Stop(ctx context.Context) error {
	return nil
}

func (srv *Server) ProtoHandler() middleware.Handler {
	return func(ctx context.Context, proto *protocol.Proto) (proto.Message, error) {
		return middleware.Chain(srv.ms)(srv.handleProto)(ctx, proto)
	}
}

func (srv *Server) handleProto(ctx context.Context, proto *protocol.Proto) (proto.Message, error) {
	return nil, nil
}
