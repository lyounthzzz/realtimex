package gateway

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/google/uuid"
	"github.com/lyounthzzz/realtimex/api/msg"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyounthzzz/realtimex/internal/gateway/operation"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport/websocket"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	srv  transport.Server
	msgc msg.MsgServiceClient
}

func NewServer() *Server {
	srv := &Server{}

	srv.srv = websocket.NewServer(
		websocket.AuthnClient(&authClient{}),
	)
	return srv
}

func (srv *Server) Start(ctx context.Context) error {
	for conn := range srv.srv.Connections() {
		srv.dispatchConn(conn)
	}

	if err := srv.srv.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (srv *Server) handleProto(ctx context.Context, proto *protocol.Proto) error {
	conn := transport.ConnFromContext(ctx)

	ctx = metadata.AppendToClientContext(ctx,
		// 连接ID
		"X-ID", uuid.New().String(),
		// 服务连接协议
		"X-Server-Protocol", "websocket",
		// 用户ID
		"X-Uid", conn.GetUser().Uid,
	)

	switch proto.Operation {
	case operation.Connect:
		if _, err := srv.msgc.Connect(ctx, &msg.ConnectReq{}); err != nil {
			_ = conn.Push(ctx, &protocol.Proto{Operation: operation.Disconnect, Body: []byte(err.Error())})
			return err
		}

	case operation.Pingreq:
		if _, err := srv.msgc.Ping(ctx, &msg.PingReq{}); err != nil {
			return err
		}
		_ = conn.Push(ctx, &protocol.Proto{Operation: operation.Pingresp})
	case operation.Publish:
		var pubReq msg.PublishReq
		_ = protojson.Unmarshal(proto.Body, &pubReq)

		if _, err := srv.msgc.Publish(ctx, &pubReq); err != nil {
			return err
		}
	case operation.Subscribe:
		var subReq msg.SubscribeReq
		_ = protojson.Unmarshal(proto.Body, &subReq)

		if _, err := srv.msgc.Subscribe(ctx, &subReq); err != nil {
			return err
		}
	}

	return nil
}

func (srv *Server) dispatchConn(conn transport.Conn) {
	user := conn.GetUser()
	if user == nil {
		_ = conn.Stop(context.Background())
		return
	}

	// 添加中间件
	conn.AddMiddleware(middleware.Logger())
	// 添加处理函数
	_ = conn.Receive(context.Background(), srv.handleProto)
	// 启动
	_ = conn.Start(context.Background())

}

func (srv *Server) Stop(ctx context.Context) error {
	return srv.srv.Stop(ctx)
}
