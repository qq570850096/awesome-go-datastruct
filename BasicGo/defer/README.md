# Defer

A focused demonstration of deferred cleanup, panic recovery, and last-in-first-out execution.

## Quick Start

```bash
go test ./BasicGo/defer
```

Run the example manually from a temporary caller if you want to inspect printed output.

## What You Will Learn

- Deferred calls run when the surrounding function returns or panics.
- Defers run in reverse registration order.
- `recover` only catches a panic from inside a deferred function on the same goroutine.
- Panic recovery is useful at process boundaries, but ordinary failures should usually be returned as errors.

## Concept Map

```text
function enters
  defer cleanup/recover
  normal work
  panic or return
deferred functions run
function exits
```

## API Surface

| Function | Purpose | Important contract |
| --- | --- | --- |
| `Error()` | Prints a start message, panics, and recovers in a deferred function. | Demonstration only; it does not return an error. |

## Guided Walkthrough

1. Read the `defer func() { ... }()` block before the panic.
2. Notice that `recover()` is checked inside the deferred function.
3. Read the test and notice that the package currently protects only "does not crash" behavior.
4. Think about how to make future examples return values so tests can assert exact behavior.

## Example Flow

```text
Start
panic(errors.New("Something Wrong!"))
defer runs
recover returns the panic value
Error returns normally
```

## Common Pitfalls

- Calling `recover` outside a deferred function and expecting it to work.
- Using panic for validation errors that callers should handle normally.
- Forgetting that deferred arguments are evaluated when the `defer` statement executes.
- Assuming recovery in one goroutine catches panics from another goroutine.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Write a function that appends `1`, `2`, `3` through defers and assert reverse order. |
| Drill | Convert `Error` into a version that returns the recovered error as a value. |
| Challenge | Add a goroutine panic example and document why the existing recover cannot catch it. |

## Quality Checklist

- Defer examples are testable without depending only on console output.
- Panic is reserved for unrecoverable or boundary-protected situations.
- Cleanup defers do not hide important returned errors.

## Related Topics

- [errors](../errors/)
- [GoRoutine](../GoRoutine/)
- [webdemo/redpacket](../../webdemo/redpacket/)
