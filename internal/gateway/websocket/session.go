package websocket

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/operation"
	"github.com/lyounthzzz/realtimex/pkg/middleware"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type sessionKey struct{}

// Session todo sync.Pool 优化
type Session struct {
	Id      string
	User    *gateway.User
	Conn    *websocket.Conn
	Server  *Server
	handler middleware.Handler
}

func (sess *Session) AddHandler(handler middleware.Handler) {
	sess.handler = handler
}

func (sess *Session) Listen() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic ", err)
		}
	}()

	log.Context(context.Background()).Infof("start listen")

	if err := sess.Connect(); err != nil {
		log.Context(context.Background()).Errorf("session connect err: ", err)
	}

	for {
		var (
			msgType int
			data    []byte
			err     error
			conn    = sess.Conn
		)

		if msgType, data, err = conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			) {
				log.Context(context.Background()).Errorf("conn read message err: ", err)
				break
			}
		}

		p := &protocol.Proto{Body: data}
		switch msgType {
		case websocket.TextMessage:
			p.Operation = operation.Publish
		case websocket.PingMessage:
			p.Operation = operation.Ping
		case websocket.PongMessage:
			p.Operation = operation.Pong
		case websocket.CloseMessage:
			break
		default:
			continue
		}

		var (
			ctx   = context.WithValue(context.Background(), sessionKey{}, sess)
			reply proto.Message
		)

		if reply, err = sess.handler(ctx, p); err != nil {
			// todo
			return
		}
		if reply != nil {
			_ = sess.WriteMessage(reply)
		}
	}

	if err := sess.Disconnect(); err != nil {
		log.Context(context.Background()).Errorf("session disconnect err: ", err)
	}
}

func (sess *Session) WriteMessage(msg proto.Message) error {
	body, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}
	return sess.Conn.WriteMessage(websocket.TextMessage, body)
}

func (sess *Session) Disconnect() error {
	_ = sess.Conn.Close()
	_, err := sess.handler(context.WithValue(context.Background(), sessionKey{}, sess), &protocol.Proto{Operation: operation.Disconnect})
	return err
}

func (sess *Session) Connect() error {
	_, err := sess.handler(context.WithValue(context.Background(), sessionKey{}, sess), &protocol.Proto{Operation: operation.Connect})
	return err
}
