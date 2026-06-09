# Goroutines

Concurrency examples covering goroutine launch, closure capture, wait groups, mutexes, cancellation, producer-consumer channels, async calls, and bounded object pools.

## Quick Start

```bash
go test ./BasicGo/GoRoutine
go test -race ./BasicGo/GoRoutine
```

Run one focused test:

```bash
go test ./BasicGo/GoRoutine -run TestNewObjPool
```

## What You Will Learn

- How goroutines run concurrently with the caller.
- Why loop variables and closures need careful handling.
- Why shared counters need synchronization.
- How `sync.WaitGroup` replaces arbitrary sleeps when waiting for work.
- How channels model asynchronous results and bounded resource pools.
- How context cancellation gives goroutines a cooperative stop signal.

## Concept Map

```text
go func()        -> concurrent work
WaitGroup        -> wait for known tasks
Mutex            -> protect shared state
channel          -> pass values / resources
context cancel   -> ask work to stop
object pool      -> bounded reusable resources
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Thread()` | Correctly captures loop variable for goroutine printing. | Demonstration uses sleep; prefer `WaitGroup` in production examples. |
| `ThreadWrong()` | Shows closure capture risk. | Intentionally flawed teaching example. |
| `CounterWrong()` | Shows unsynchronized shared mutation. | Intentionally race-prone. |
| `Counter()` | Protects a counter with `sync.Mutex`. | Uses sleep; `WaitGroupExam` is more deterministic. |
| `WaitGroupExam()` | Counts with `WaitGroup` and mutex. | Waits for all goroutines before reading. |
| `AsnyService()` | Starts async work and returns a result channel. | Channel is buffered with capacity 1. |
| `Producer`, `Consumer` | Producer-consumer pattern. | Producer closes the channel. |
| `Cancel()` | Cancels goroutines through context. | Goroutines poll for cancellation. |
| `ObjPool` | Bounded reusable object pool. | `GetObj` can time out; `ReleaseObj` can overflow. |

## Guided Walkthrough

1. Compare `Thread` and `ThreadWrong`.
2. Compare `CounterWrong`, `Counter`, and `WaitGroupExam`.
3. Read `AsnyService` and trace when the result channel receives a value.
4. Read `Producer` and `Consumer`; identify who owns channel closing.
5. Finish with `ObjPool`, which is a practical channel-as-resource-pool example.

## Example

```go
pool := NewObjPool(2)
obj, err := pool.GetObj(10 * time.Millisecond)
if err != nil {
    return err
}
defer pool.ReleaseObj(obj)
```

The channel capacity is the resource limit. Borrowers block or time out when the pool is empty.

## Common Pitfalls

- Reading shared state before all goroutines finish.
- Relying on `time.Sleep` as synchronization.
- Calling `WaitGroup.Done` while still holding a lock if later code might block.
- Releasing the same pooled object twice.
- Assuming cancellation interrupts work automatically; code must observe the signal.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Rewrite `Counter` to use `WaitGroup` instead of sleep. |
| Drill | Add a timeout test for `ObjPool.GetObj` when the pool is empty. |
| Challenge | Replace one polling loop with a blocking receive on `ctx.Done()`. |

## Quality Checklist

- Race detector passes.
- Every goroutine has a completion path.
- Channel closing has a single owner.
- Tests avoid timing assumptions where possible.

## Related Topics

- [channelselect](../channelselect/)
- [sharedvars](../sharedvars/)
- [context](../context/)
