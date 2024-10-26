package ctxutil

import (
	"context"
	"testing"
	"time"
)

func TestSleepContext(t *testing.T) {
	tests := []struct {
		name     string
		timeout  time.Duration
		sleep    time.Duration
		wantWait bool
	}{
		{
			name:     "normal sleep",
			timeout:  time.Second,
			sleep:    time.Millisecond * 100,
			wantWait: true,
		},
		{
			name:     "context canceled",
			timeout:  time.Millisecond * 100,
			sleep:    time.Second,
			wantWait: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			start := time.Now()
			SleepContext(ctx, tt.sleep)
			elapsed := time.Since(start)

			if tt.wantWait && elapsed < tt.sleep {
				t.Errorf("SleepContext() returned too early, elapsed: %v, want: %v", elapsed, tt.sleep)
			}
			if !tt.wantWait && elapsed >= tt.sleep {
				t.Errorf("SleepContext() didn't cancel in time, elapsed: %v, timeout: %v", elapsed, tt.timeout)
			}
		})
	}
}

func TestDelayContext(t *testing.T) {
	tests := []struct {
		name    string
		timeout time.Duration
		delay   time.Duration
		wantErr bool
	}{
		{
			name:    "context canceled",
			timeout: time.Millisecond * 100,
			delay:   time.Millisecond * 10,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			err := DelayContext(ctx, tt.delay)
			if (err != nil) != tt.wantErr {
				t.Errorf("DelayContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != context.DeadlineExceeded {
				t.Errorf("DelayContext() error = %v, want context.DeadlineExceeded", err)
			}
		})
	}
}

func TestCopyContextWithOutTimeOut(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  interface{}
	}{
		{
			name: "copy context values",
			key:  "test-key",
			val:  "test-value",
		},
	}
	type Tmp string
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create context with timeout and value
			ctx := context.WithValue(context.Background(), Tmp(tt.key), tt.val)
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			// Copy context without timeout
			newCtx := CopyContextWithOutTimeOut(ctx)

			// Check if deadline is removed
			if deadline, ok := newCtx.Deadline(); ok {
				t.Errorf("CopyContextWithOutTimeOut() deadline = %v, want no deadline", deadline)
			}

			// Check if value is preserved
			if got := newCtx.Value(Tmp(tt.key)); got != tt.val {
				t.Errorf("CopyContextWithOutTimeOut() value = %v, want %v", got, tt.val)
			}

			// Check Done() channel
			if newCtx.Done() != nil {
				t.Error("CopyContextWithOutTimeOut() Done() should return nil")
			}

			// Check Err()
			if err := newCtx.Err(); err != nil {
				t.Errorf("CopyContextWithOutTimeOut() Err() = %v, want nil", err)
			}
		})
	}
}
