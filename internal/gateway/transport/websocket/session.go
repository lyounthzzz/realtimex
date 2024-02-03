package websocket

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyounthzzz/realtimex/internal/gateway/operation"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

type sessionKey struct{}

var _ transport.Conn = (*Session)(nil)

// Session todo sync.Pool 优化
type Session struct {
	User    *gateway.User
	Conn    *websocket.Conn
	Server  *Server
	sendBuf chan *protocol.Proto
	handle  middleware.Handler
	ms      []middleware.Middleware
}

func (sess *Session) Start(ctx context.Context) error {
	go sess.read()
	go sess.write()
	return nil
}

func (sess *Session) read() {
	for {
		var (
			msgType int
			data    []byte
			err     error
		)

		if msgType, data, err = sess.Conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			) {
				log.Context(context.Background()).Errorf("sess read message err: ", err)
				break
			}
		}

		p := &protocol.Proto{Body: data}
		switch msgType {
		case websocket.TextMessage:
			p.Operation = operation.Publish
		case websocket.PingMessage:
			p.Operation = operation.Pingreq
		case websocket.CloseMessage:
			p.Operation = operation.Disconnect
			break
		default:
			continue
		}

		if sess.handle == nil {
			continue
		}

		ctx := transport.ContextWithConn(context.Background(), sess)
		if err = middleware.Chain(sess.ms...)(sess.handle)(ctx, p); err != nil {
			continue
		}
	}
}

func (sess *Session) write() {
	for p := range sess.sendBuf {
		body, err := protojson.Marshal(p)
		if err != nil {
			continue
		}
		_ = sess.Conn.WriteMessage(websocket.TextMessage, body)
	}
}

func (sess *Session) Receive(ctx context.Context, handle middleware.Handler) error {
	if sess.handle != nil {
		sess.handle = handle
	}
	return nil
}

func (sess *Session) Push(ctx context.Context, proto *protocol.Proto) error {
	sess.sendBuf <- proto
	return nil
}

func (sess *Session) GetUser() *gateway.User {
	return sess.User
}

func (sess *Session) Stop(ctx context.Context) error {
	close(sess.sendBuf)
	_ = sess.Conn.Close()
	return nil
}

func (sess *Session) AddMiddleware(ms ...middleware.Middleware) {
	sess.ms = ms
}
