# Context

## Goal

Use `context.Context` for deadlines, cancellation, and request-scoped values without turning it into a general parameter bag.

## Key Ideas

- A derived context should be canceled by the owner that creates it.
- Deadlines and timeouts are cooperative; work must observe `ctx.Done()`.
- Request values should use private key types to avoid collisions.
- A canceled context returns a stable reason through `ctx.Err()`.

## Repository Code Map

| File | What to read for |
| --- | --- |
| timeout.go | Request ID helpers, timeout wrapping, and cancelable work examples. |
| timeout_test.go | Timeout and value-propagation tests. |

## Core Invariant

Long-running work should either complete normally or return when the context is canceled.

## Practice Tasks

- Add a test where work finishes before the timeout.
- Add a test that checks the exact cancellation error.
- Pass a request ID through a nested context and read it back.

## Test Command

```bash
go test ./BasicGo/context
```

## Related Topics

- [channelselect](../channelselect/)
- [webdemo/http_basic](../../webdemo/http_basic/)
