package websocket

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/pkg/processor"
	"net/http"
	"time"
)

const (
	defaultAddress          = ":8000"
	defaultPath             = "/gateway/ws/v1"
	defaultSessConnectedCap = 2048
	defaultSessDroppedCap   = 2048
)

var defaultWebsocketUpgrader = &websocket.Upgrader{
	HandshakeTimeout:  time.Second,
	ReadBufferSize:    10 * 1024 * 1024, // 10k
	WriteBufferSize:   64 * 1024 * 1024, // 64k
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: true,
}

type Server struct {
	*http.Server
	address           string
	path              string
	router            *mux.Router
	upgrader          *websocket.Upgrader
	sessionsConnected chan *Session
	sessionsToDropped chan *Session
	authnC            gateway.AuthnServiceClient
	protoHandler      processor.Handler
}

func NewServer(opts ...Option) *Server {
	srv := &Server{
		address:           defaultAddress,
		path:              defaultPath,
		router:            mux.NewRouter(),
		upgrader:          defaultWebsocketUpgrader,
		sessionsConnected: make(chan *Session, defaultSessConnectedCap),
		sessionsToDropped: make(chan *Session, defaultSessDroppedCap),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (srv *Server) Start(ctx context.Context) error {
	srv.initRouter()

	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (srv *Server) Stop(ctx context.Context) error {
	_ = srv.Server.Shutdown(ctx)
	close(srv.sessionsConnected)
	close(srv.sessionsToDropped)
	return nil
}

func (srv *Server) initRouter() {
	srv.router.Handle(srv.path, srv.websocketHandler())
}

func (srv *Server) websocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var (
			ctx   = req.Context()
			token = req.Header.Get("Authorization")
			user  *gateway.User
			conn  *websocket.Conn
			err   error
		)

		if conn, err = srv.upgrader.Upgrade(w, req, nil); err != nil {
			return
		}
		// set close ping pong handler
		// 用户鉴权
		if user, err = srv.authn(ctx, token); err != nil {
			return
		}

		session := &Session{
			Id:     "xx",
			User:   user,
			Conn:   conn,
			Server: srv,
		}
		srv.sessionsConnected <- session

		go session.Listen()
	})
}

func (srv *Server) authn(ctx context.Context, token string) (*gateway.User, error) {
	return srv.authnC.DoAuthn(ctx, &gateway.DoAuthNRequest{Token: token})
}

func (srv *Server) Sessions() chan *Session {
	return srv.sessionsConnected
}

func (srv *Server) CloseSession(sess *Session) error {
	srv.sessionsToDropped <- sess
	return nil
}

func (srv *Server) closeSessionBackground() {
	for sess := range srv.sessionsToDropped {
		_ = sess.Conn.Close()
	}
}
