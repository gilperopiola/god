package god

import (
	"context"
	"time"
)

type Ctx = context.Context

func NewCtx() Ctx {
	return context.Background()
}

func NewCtxWithTimeout(duration time.Duration) (Ctx, context.CancelFunc) {
	return context.WithTimeout(NewCtx(), duration)
}
