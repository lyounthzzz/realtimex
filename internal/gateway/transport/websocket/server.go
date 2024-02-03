package websocket

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/lyounthzzz/realtimex/internal/gateway/transport"
	"net/http"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	*http.Server
	address  string
	path     string
	router   *mux.Router
	upgrader *websocket.Upgrader
	sessions chan transport.Conn
	authnC   gateway.AuthnServiceClient
	receiver transport.Receiver
}

func (srv *Server) Connections() chan transport.Conn {
	return srv.sessions
}

func NewServer(opts ...Option) *Server {
	srv := &Server{
		address:  defaultAddress,
		path:     defaultPath,
		router:   mux.NewRouter(),
		upgrader: defaultWebsocketUpgrader,
		sessions: make(chan transport.Conn, defaultSessConnectedCap),
	}

	for _, opt := range opts {
		opt(srv)
	}
	srv.Server = &http.Server{
		Addr:    srv.address,
		Handler: srv.router,
	}

	return srv
}

func (srv *Server) Start(ctx context.Context) error {
	srv.router.Handle(srv.path, srv.websocketHandler())

	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
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
		// set close ping pong handle
		// 用户鉴权
		if user, err = srv.authn(ctx, token); err != nil {
			return
		}

		session := &Session{
			User:   user,
			Conn:   conn,
			Server: srv,
		}
		srv.sessions <- session
	})
}

func (srv *Server) authn(ctx context.Context, token string) (*gateway.User, error) {
	return srv.authnC.DoAuthn(ctx, &gateway.DoAuthNRequest{Token: token})
}

func (srv *Server) Stop(ctx context.Context) error {
	_ = srv.Server.Shutdown(ctx)
	close(srv.sessions)
	return nil
}
