package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"google.golang.org/protobuf/encoding/protojson"
)

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

		switch msgType {
		case websocket.TextMessage:
			proto := &protocol.Proto{}

			if err = protojson.Unmarshal(data, proto); err != nil {
				// todo
				return
			}
			if err = sess.Server.protoHandler(context.Background(), proto); err != nil {
				// todo
				return
			}
		case websocket.CloseMessage:
			sess.Close()
		case websocket.PingMessage:
			// todo
		case websocket.PongMessage:
			// todo
		default:

		}
	}
}

func (sess *Session) Write(proto *protocol.Proto) error {
	body, err := protojson.Marshal(proto)
	if err != nil {
		return err
	}
	return sess.Conn.WriteMessage(websocket.TextMessage, body)
}

func (sess *Session) Close() {
	sess.Server.sessionsToDropped <- sess
}
