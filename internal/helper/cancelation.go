package helper

import (
	"context"
)

// CheckDeadline check if context has cancelled
func CheckDeadline(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
