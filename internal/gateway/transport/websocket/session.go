package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"github.com/lyounthzzz/realtimex/internal/gateway/operation"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type sessionKey struct{}

// Session todo sync.Pool 优化
type Session struct {
	Id     string
	User   *gateway.User
	Conn   *websocket.Conn
	Server *Server
}

func (sess *Session) Listen() {
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
				// todo
				return
			}
		}

		var p *protocol.Proto
		switch msgType {
		case websocket.TextMessage:
			p = &protocol.Proto{}
			if err = protojson.Unmarshal(data, p); err != nil {
				// todo
				continue
			}
		case websocket.CloseMessage:
			p = &protocol.Proto{
				Operation: operation.Close,
			}
		case websocket.PingMessage:
			p = &protocol.Proto{
				Operation: operation.Ping,
			}
		case websocket.PongMessage:
			p = &protocol.Proto{
				Operation: operation.Pong,
			}
		default:
			continue
		}

		var (
			ctx   = context.WithValue(context.Background(), sessionKey{}, sess)
			reply proto.Message
		)
		if reply, err = sess.Server.protoHandler(ctx, p); err != nil {
			// todo
			return
		}
		if reply != nil {
			_ = sess.WriteMessage(reply)
		}
	}
}

func (sess *Session) WriteMessage(msg proto.Message) error {
	body, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}
	return sess.Conn.WriteMessage(websocket.TextMessage, body)
}

func (sess *Session) Close() {
	sess.Server.sessionsToDropped <- sess
}
