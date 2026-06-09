# Channels And Select

Coordination examples for fan-in, tagged aggregation, ticking work, draining streams, and cooperative cancellation.

## Quick Start

```bash
go test ./BasicGo/channelselect
go test ./BasicGo/channelselect -run TestFanIn
```

## What You Will Learn

- How channels transfer values and ownership between goroutines.
- How `select` waits on input, output, timers, and cancellation signals.
- How to close output channels after all workers finish.
- Why context-aware goroutines are easier to stop cleanly.
- How tests can avoid sleeps by controlling channels directly.

## Concept Map

```text
input channel(s)
      |
      v
worker goroutines -- ctx.Done() --> stop
      |
      v
output channel -- close --> consumer finishes
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `FanIn(ctx, inputs...)` | Merge string streams into one output stream. | Ignores nil inputs and closes output when workers finish. |
| `TaggedMessage` | Preserve stream source with each log payload. | `Source` is the map key from `AggregateLogs`. |
| `AggregateLogs(ctx, streams)` | Merge named streams into tagged messages. | Output ordering is concurrent and not globally stable. |
| `Ticker(ctx, interval)` | Emit time values until cancellation. | Stops the underlying ticker before exiting. |
| `Drain(ctx, in)` | Collect integers until input closes or context cancels. | May return a partial slice on cancellation. |

## Guided Walkthrough

1. Read `FanIn` first. It shows the standard pattern: one goroutine per input, one `WaitGroup`, one closer goroutine.
2. Read `AggregateLogs` next. It adds metadata while keeping the same lifecycle.
3. Read `Ticker` and focus on `defer t.Stop()`.
4. Finish with `Drain`; it is intentionally simple, but it makes cancellation observable in tests.

## Example

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

out := FanIn(ctx, serviceA, serviceB)
for msg := range out {
    fmt.Println(msg)
}
```

The consumer ranges until `FanIn` closes the output. If `ctx` is canceled, workers return and the closer goroutine closes `out`.

## Common Pitfalls

- Closing a channel from the receiver side.
- Forgetting to close the output channel after fan-in workers exit.
- Assuming concurrent fan-in preserves source ordering.
- Sending to an output channel without also listening for `ctx.Done()`.
- Using arbitrary sleeps in tests instead of explicit channel control.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a `FanIn` test with one nil input channel. |
| Drill | Add a cancellation test where `Drain` returns a partial result. |
| Challenge | Implement `FanInLimit` that accepts a maximum number of active input workers. |

## Quality Checklist

- Every goroutine has a path to exit.
- The producer side owns channel closing.
- Tests cover nil channels, early cancellation, and normal completion.
- Race detector stays clean: `go test -race ./BasicGo/channelselect`.

## Related Topics

- [GoRoutine](../GoRoutine/)
- [context](../context/)
- [sharedvars](../sharedvars/)
