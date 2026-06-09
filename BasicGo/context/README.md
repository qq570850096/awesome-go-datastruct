# Context

Timeout, cancellation, and request value examples built around `context.Context`.

## Quick Start

```bash
go test ./BasicGo/context
go test ./BasicGo/context -run TestFetchWithTimeout
```

## What You Will Learn

- How to derive a context with a timeout and release its timer.
- How cooperative work observes `ctx.Done()`.
- How private key types prevent request-value collisions.
- How cancellation errors travel back to callers.
- Why context should be passed explicitly and usually first.

## Concept Map

```text
parent context
      |
      +-- WithTimeout -> child context -- ctx.Done() --> work stops
      |
      +-- WithValue   -> request metadata -> read by downstream code
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `ContextWithRequestID(parent, id)` | Attach a request ID using a private key type. | Does not mutate the parent context. |
| `RequestIDFromContext(ctx)` | Read the request ID if present. | Returns an empty string when absent or wrong type. |
| `FetchWithTimeout(parent, timeout, work)` | Run work under a timeout. | Cancels the child context before returning. |
| `ExampleSlowWork(delay, response)` | Build sample work that respects cancellation. | Returns `ctx.Err()` when canceled. |
| `ExampleCancelableWork(ctx)` | Minimal function that exits on context cancellation. | Useful as a test fixture. |

## Guided Walkthrough

1. Start with `ContextWithRequestID` and notice the unexported key type.
2. Read `FetchWithTimeout`; the `defer cancel()` is the resource cleanup.
3. Read `ExampleSlowWork`; it is a model for work that can either finish or stop.
4. Inspect `timeout_test.go` and compare successful work with timed-out work.

## Example

```go
result, err := FetchWithTimeout(context.Background(), 20*time.Millisecond, ExampleSlowWork(time.Second, "ok"))
fmt.Println(result, err) // "", context deadline exceeded
```

The timeout does not kill a goroutine by force. It only signals through the context; the work must listen.

## Common Pitfalls

- Forgetting to call the `cancel` function returned by `WithTimeout`.
- Storing ordinary function parameters in context values.
- Using exported string keys and colliding with another package.
- Ignoring `ctx.Err()` and returning an unrelated error.
- Passing `nil` instead of `context.Background()` or `context.TODO()`.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a test where work finishes before the timeout. |
| Drill | Add nested request-ID propagation and read it at the deepest layer. |
| Challenge | Implement a retry helper that stops retrying when the context is canceled. |

## Quality Checklist

- Context is the first argument on context-aware functions.
- Derived contexts are canceled by their owner.
- Tests cover success, deadline, and explicit cancellation.
- Request values are small, request-scoped, and typed.

## Related Topics

- [channelselect](../channelselect/)
- [webdemo/http_basic](../../webdemo/http_basic/)
- [sharedvars](../sharedvars/)
