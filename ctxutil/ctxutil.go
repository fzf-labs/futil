package ctxutil

import (
	"context"
	"time"
)

// SleepContext sleep with context
func SleepContext(ctx context.Context, delay time.Duration) {
	timer := time.NewTimer(delay)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
	}
}

// DelayContext delay with context
func DelayContext(ctx context.Context, delay time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(delay)
		}
	}
}

// CopyContextWithOutTimeOut copy context without timeout
func CopyContextWithOutTimeOut(ctx context.Context) context.Context {
	return contextWithOutTimeOut{ctx}
}

type contextWithOutTimeOut struct {
	ctx context.Context
}

func (contextWithOutTimeOut) Deadline() (time.Time, bool) { return time.Time{}, false }
func (contextWithOutTimeOut) Done() <-chan struct{}       { return nil }
func (contextWithOutTimeOut) Err() error                  { return nil }
func (x contextWithOutTimeOut) Value(key interface{}) interface{} {
	return x.ctx.Value(key)
}
