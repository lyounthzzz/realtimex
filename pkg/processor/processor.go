package processor

import (
	"context"
	"github.com/lyounthzzz/realtimex/api/protocol"
)

type Handler func(ctx context.Context, proto *protocol.Proto) error

type Processor func(Handler) Handler

func Chain(ps ...Processor) Processor {
	return func(next Handler) Handler {
		for i := len(ps) - 1; i >= 0; i-- {
			next = ps[i](next)
		}
		return next
	}
}
