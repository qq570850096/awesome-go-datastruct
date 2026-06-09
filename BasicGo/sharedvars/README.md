# Shared Variables

Race-free access to shared mutable state with mutexes, atomics, and one-time initialization.

## Quick Start

```bash
go test ./BasicGo/sharedvars
go test -race ./BasicGo/sharedvars
```

## What You Will Learn

- How a mutex protects a critical section.
- How atomic counters work for simple numeric state.
- How `sync.Once` guarantees one-time initialization.
- How tests and the race detector complement each other.

## Concept Map

```text
shared state
      |
      +-- mutex  -> lock, mutate/read, unlock
      +-- atomic -> single low-level operation
      +-- once   -> initialize exactly one time
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Counter` | Mutex-protected integer counter. | All reads and writes use the same mutex. |
| `Counter.Inc()` | Increment safely. | Mutates under lock. |
| `Counter.Value()` | Read safely. | Reads under lock. |
| `AtomicCounter` | Atomic integer counter. | Uses atomic operations consistently. |
| `OnceValue` | Lazily initialize one string value. | First initialization wins. |

## Guided Walkthrough

1. Read `Counter` first; it is the clearest critical-section example.
2. Read `AtomicCounter` and compare the smaller API with the mutex version.
3. Read `OnceValue` and notice that callers can race to initialize but only one function runs.
4. Inspect tests and count how many goroutines touch each value.

## Example

```go
var c Counter
c.Inc()
fmt.Println(c.Value()) // 1
```

For a compound invariant, prefer a mutex. Atomics are best when the operation is simple and the invariant is narrow.

## Common Pitfalls

- Protecting writes with a mutex but reading without it.
- Mixing atomic and non-atomic access to the same variable.
- Using `sync.Once` when the value must be reset.
- Treating `-race` as a replacement for behavior tests.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a test that reads `Counter.Value` while goroutines increment. |
| Drill | Add benchmarks comparing mutex and atomic counters. |
| Challenge | Implement a concurrent-safe map wrapper with clear locking boundaries. |

## Quality Checklist

- Race detector passes.
- State ownership is obvious from the type.
- Locks are held for the shortest practical critical section.
- Tests assert final values after all goroutines finish.

## Related Topics

- [GoRoutine](../GoRoutine/)
- [channelselect](../channelselect/)
- [context](../context/)
