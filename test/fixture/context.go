package fixture

import (
	"context"
	"time"
)

// CtxEnded creates dummy context with cancelled state.
func CtxEnded() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))
	defer cancel()
	return ctx
}
