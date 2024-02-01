package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lyounthzzz/realtimex/api/protocol"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"time"
)

func Logger() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, p *protocol.Proto) (proto.Message, error) {
			var (
				op        = p.Operation
				ingress   = len(p.Body)
				startTime = time.Now()
				reply     proto.Message
				replyBody []byte
				err       error
			)

			reply, err = next(ctx, p)
			if err == nil {
				replyBody, _ = protojson.Marshal(p)
			}

			if err != nil {
				log.Context(ctx).Errorf(err.Error(), op, ingress, time.Since(startTime).Seconds())
			} else {
				log.Context(ctx).Infof(string(replyBody), op, ingress, time.Since(startTime).Seconds())
			}

			return reply, err
		}
	}
}
