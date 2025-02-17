package ctxutils

import (
	"context"
	"time"

	internaltime "github.com/aereal/ctxutils/internal/time"
)

// ContextWithNarrowedDeadline returns a new context with a shortened deadline
// applying a grace period reduction.
func ContextWithNarrowedDeadline(ctx context.Context, grace time.Duration) (context.Context, func()) {
	return ContextWithNarrowedDeadlineCause(ctx, grace, nil)
}

// ContextWithNarrowedDeadlineCause returns a new context with a shortened deadline
// and an associated cause error, applying a grace period reduction.
func ContextWithNarrowedDeadlineCause(ctx context.Context, grace time.Duration, cause error) (context.Context, func()) {
	deadline, _ := ctx.Deadline()
	return context.WithDeadlineCause(ctx, narrowDeadline(deadline, grace), cause)
}

func narrowDeadline(deadline time.Time, grace time.Duration) (narrowed time.Time) {
	now := internaltime.Now()
	if deadline.IsZero() { // zero deadline value implies no deadline bound for the context
		return now.Add(grace)
	}
	remaining := deadline.Sub(now)
	if remaining < grace {
		return deadline
	}
	return deadline.Add(grace * -1)
}
