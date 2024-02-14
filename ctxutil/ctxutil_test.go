package ctxutil

import (
	"context"
	"testing"
	"time"
)

func TestSleepContext(t *testing.T) {
	type args struct {
		ctx   context.Context
		delay time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SleepContext(tt.args.ctx, tt.args.delay)
		})
	}
}
