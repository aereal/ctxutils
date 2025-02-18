package ctxutils_test

import (
	"context"
	"fmt"
	"time"

	"github.com/aereal/ctxutils"
	internaltime "github.com/aereal/ctxutils/internal/time"
)

func ExampleContextWithNarrowedDeadline() {
	base, err := time.ParseInLocation(time.DateTime, "2006-01-02 15:04:05", time.UTC)
	if err != nil {
		panic(err)
	}
	internaltime.SetNow(base)

	originalDeadline := base.Add(time.Hour * 1)
	ctx, _ := context.WithDeadline(context.Background(), originalDeadline) //nolint:govet
	narrowedCtx, _ := ctxutils.ContextWithNarrowedDeadline(ctx, time.Minute*4+time.Second*5)
	narrowedDeadline, ok := narrowedCtx.Deadline()
	if !ok {
		panic("deadline is not narrowed")
	}
	fmt.Printf("original deadline: %s\n", originalDeadline.In(time.UTC).Format(time.DateTime))
	fmt.Printf("new deadline: %s\n", narrowedDeadline.In(time.UTC).Format(time.DateTime))
	// Output:
	// original deadline: 2006-01-02 16:04:05
	// new deadline: 2006-01-02 16:00:00
}
