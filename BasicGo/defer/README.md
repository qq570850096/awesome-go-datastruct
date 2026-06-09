# Defer

## Goal

Study when deferred functions run and how they interact with panic-style control flow.

## Key Ideas

- Deferred calls run after the surrounding function begins returning.
- Multiple defers run in last-in, first-out order.
- `recover` only works inside a deferred function on the same goroutine.
- Defer is useful for cleanup, but ordinary errors should still be returned explicitly.

## Repository Code Map

| File | What to read for |
| --- | --- |
| defer.go | A small defer, panic, and recover demonstration. |
| defer_test.go | Smoke test for the example function. |

## Core Invariant

Cleanup code registered with `defer` should run exactly once for the function activation that registered it.

## Practice Tasks

- Add a testable function that records defer order in a slice.
- Convert a panic example into an explicit error-returning function.
- Explain why recovering in a different goroutine does not catch the panic.

## Test Command

```bash
go test ./BasicGo/defer
```

## Related Topics

- [errors](../errors/)
- [GoRoutine](../GoRoutine/)
