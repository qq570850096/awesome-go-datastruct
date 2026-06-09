# Goroutines

## Goal

Explore goroutine scheduling, wait groups, mutex-protected counters, cancellation checks, producer-consumer channels, and a small reusable object pool.

## Key Ideas

- A goroutine starts work concurrently, but completion still needs coordination.
- `sync.WaitGroup` waits for a known set of tasks.
- Shared state needs synchronization when multiple goroutines mutate it.
- Channels can model producer-consumer handoff and bounded resource pools.

## Repository Code Map

| File | What to read for |
| --- | --- |
| Goroutine.go | Goroutine examples, wait groups, cancellation, producer-consumer flow, and object pool. |
| Goroutine_test.go | Behavioral tests for thread examples, cancellation, and pool creation. |

## Core Invariants

- Every `WaitGroup.Add` should have a matching `Done`.
- Shared counters should not be updated concurrently without synchronization.
- A borrowed object should be returned to the pool at most once.

## Practice Tasks

- Compare `CounterWrong` and `Counter` under `go test -race`.
- Add a timeout test for `ObjPool.GetObj`.
- Rewrite one cancellation check to use `context.Context`.

## Test Command

```bash
go test ./BasicGo/GoRoutine
go test -race ./BasicGo/GoRoutine
```

## Related Topics

- [channelselect](../channelselect/)
- [sharedvars](../sharedvars/)
