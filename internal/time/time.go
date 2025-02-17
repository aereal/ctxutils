package time

import (
	"sync/atomic"
	"testing"
	"time"
)

var now = new(atomic.Pointer[time.Time])

func SetNowForTest(t *testing.T, tt time.Time) {
	t.Helper()
	cleanup := SetNow(tt)
	t.Cleanup(cleanup)
}

func SetNow(t time.Time) func() {
	old := now.Swap(&t)
	return func() {
		now.Store(old)
	}
}

func Now() time.Time {
	if fixed := now.Load(); fixed != nil {
		return *fixed
	}
	return time.Now()
}

func Since(t time.Time) time.Duration { return Now().Sub(t) }

func Until(t time.Time) time.Duration { return t.Sub(Now()) }
