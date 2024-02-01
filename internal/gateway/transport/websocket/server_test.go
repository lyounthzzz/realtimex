package websocket

import (
	"context"
	"github.com/lyounthzzz/realtimex/api/gateway"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

type authClient struct {
	gateway.UnimplementedAuthnServiceServer
}

func (auth *authClient) DoAuthn(ctx context.Context, in *gateway.DoAuthNRequest, opts ...grpc.CallOption) (*gateway.User, error) {
	return &gateway.User{Uid: "1"}, nil
}

func TestWebsocketServer(t *testing.T) {
	srv := NewServer(AuthnClient(&authClient{}))

	err := srv.Start(context.Background())
	require.NoError(t, err)
}
