package ctxutils_test

import (
	"context"
	"testing"
	"time"

	"github.com/aereal/ctxutils"
	internaltime "github.com/aereal/ctxutils/internal/time"
)

func TestContextWithNarrowedDeadline(t *testing.T) {
	now := time.Now().Truncate(0)
	testCases := []struct {
		want             time.Time
		originalDeadline time.Time
		name             string
		grace            time.Duration
	}{
		{
			name:             "ok",
			originalDeadline: now.Add(time.Minute * 3),
			grace:            time.Minute * 1,
			want:             now.Add(time.Minute * 2), // 3 - 1 = 2
		},
		{
			name:             "no grace remained",
			originalDeadline: now.Add(time.Minute * 3),
			grace:            time.Minute * 4,
			want:             now.Add(time.Minute * 3),
		},
		{
			name:             "no deadline bound",
			originalDeadline: time.Time{},
			grace:            time.Minute * 1,
			want:             now.Add(time.Minute * 1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			internaltime.SetNowForTest(t, now)
			ctx := context.Background() //nolint:usetesting
			if !tc.originalDeadline.IsZero() {
				ctx, _ = context.WithDeadline(ctx, tc.originalDeadline) //nolint:govet
			}
			gotCtx, _ := ctxutils.ContextWithNarrowedDeadline(ctx, tc.grace)
			gotDeadline, hasDeadline := gotCtx.Deadline()
			if deadlineExpected := !tc.want.IsZero(); deadlineExpected != hasDeadline {
				t.Errorf("expected deadline (%v) but deadline has set? (%v)", deadlineExpected, hasDeadline)
			}
			if !tc.want.Equal(gotDeadline) {
				t.Errorf("deadline mismatch:\n\twant: %s\n\t got: %s", tc.want, gotDeadline)
			}
		})
	}
}
