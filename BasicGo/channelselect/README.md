# Channels And Select

## Goal

Practice channel coordination with fan-in, tagged aggregation, tickers, cancellation, and deterministic draining.

## Key Ideas

- Channels carry ownership of values between goroutines.
- `select` lets one goroutine wait on several communication events.
- Context cancellation should stop background senders and receivers.
- Closing an output channel tells downstream consumers that no more values are coming.

## Repository Code Map

| File | What to read for |
| --- | --- |
| select.go | `FanIn`, `AggregateLogs`, `Ticker`, and `Drain`. |
| select_test.go | Cancellation, ticker, and aggregation tests. |

## Core Invariants

- Goroutines must exit when the provided context is canceled.
- Output channels must be closed by the producer side.
- Fan-in should preserve values, not a global ordering guarantee.

## Practice Tasks

- Add a test for canceling `FanIn` before any input sends.
- Add a stream name collision case for `AggregateLogs`.
- Explain why `Drain` may return early when the context is canceled.

## Test Command

```bash
go test ./BasicGo/channelselect
```

## Related Topics

- [GoRoutine](../GoRoutine/)
- [context](../context/)
