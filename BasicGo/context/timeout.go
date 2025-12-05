package contextdemo

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type requestIDKey struct{}

// ContextWithRequestID 将 request id 塞进 context，演示 WithValue 的典型姿势。
func ContextWithRequestID(parent context.Context, id string) context.Context {
	return context.WithValue(parent, requestIDKey{}, id)
}

// RequestIDFromContext 读取 request id，提高日志可观测性。
func RequestIDFromContext(ctx context.Context) string {
	if v, ok := ctx.Value(requestIDKey{}).(string); ok {
		return v
	}
	return ""
}

// FetchWithTimeout 在指定超时内执行 work，超时返回 context.DeadlineExceeded。
func FetchWithTimeout(parent context.Context, timeout time.Duration, work func(context.Context) (string, error)) (string, error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	type result struct {
		data string
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		val, err := work(ctx)
		ch <- result{data: val, err: err}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res.data, res.err
	}
}

// ExampleSlowWork 用来模拟耗时请求，可在测试中注入。
func ExampleSlowWork(delay time.Duration, response string) func(context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		select {
		case <-time.After(delay):
			return fmt.Sprintf("data:%s", response), nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// ExampleCancelableWork 用来展示 ctx.Done 的典型用法。
func ExampleCancelableWork(ctx context.Context) error {
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if RequestIDFromContext(ctx) != "" {
				return nil
			}
		case <-time.After(20 * time.Millisecond):
			return errors.New("no request id received in time")
		}
	}
}
