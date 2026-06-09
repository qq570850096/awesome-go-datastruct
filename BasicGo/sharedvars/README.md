# Shared Variables

## Goal

Practice safe access to shared mutable state with mutexes, atomics, and one-time initialization.

## Key Ideas

- A mutex protects a critical section around shared state.
- Atomic operations are useful for small numeric state with simple invariants.
- `sync.Once` guarantees one-time initialization even under concurrency.
- Race-free behavior should be validated with tests and the race detector.

## Repository Code Map

| File | What to read for |
| --- | --- |
| sharedvars.go | Mutex counter, atomic counter, and `sync.Once` value. |
| sharedvars_test.go | Concurrent correctness tests. |

## Core Invariants

- Every read and write of mutex-protected state should use the same mutex.
- Atomic state should not be mixed with unsynchronized ordinary access.
- Initialization guarded by `sync.Once` should be idempotent from the caller's view.

## Practice Tasks

- Run the package with `-race` and inspect the result.
- Add a reset-free benchmark for `Counter` and `AtomicCounter`.
- Explain when a mutex is clearer than atomics.

## Test Command

```bash
go test ./BasicGo/sharedvars
go test -race ./BasicGo/sharedvars
```

## Related Topics

- [GoRoutine](../GoRoutine/)
- [channelselect](../channelselect/)
