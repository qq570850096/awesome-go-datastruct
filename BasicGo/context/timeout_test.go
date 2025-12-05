package contextdemo

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchWithTimeout(t *testing.T) {
	ctx := context.Background()
	data, err := FetchWithTimeout(ctx, 50*time.Millisecond, ExampleSlowWork(10*time.Millisecond, "ok"))
	if err != nil || data != "data:ok" {
		t.Fatalf("unexpected result data=%s err=%v", data, err)
	}

	_, err = FetchWithTimeout(ctx, 5*time.Millisecond, ExampleSlowWork(50*time.Millisecond, "slow"))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}
}

func TestRequestIDPropagation(t *testing.T) {
	ctx := ContextWithRequestID(context.Background(), "req-42")
	child, cancel := context.WithCancel(ctx)
	defer cancel()

	if got := RequestIDFromContext(child); got != "req-42" {
		t.Fatalf("expected request id req-42, got %s", got)
	}
	if err := ExampleCancelableWork(child); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
