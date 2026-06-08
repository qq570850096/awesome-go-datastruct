# Queue

## Goal

Learn first-in first-out behavior and the role queues play in BFS and scheduling.

## Prerequisites

Slices and basic tests.

## Core Invariant

The earliest enqueued item is the next item dequeued.

## Complexity

Enqueue: Time O(1) amortized. Dequeue: depends on implementation, typically O(1) with a head index. Space O(n).

## Practice Tasks

- Add an empty-dequeue test.
- Add a long enqueue/dequeue sequence.
- Use the queue to trace BFS manually.

## Test Command

```bash
go test ./queue
```

## Related Topics

- [stack](../stack/)
- [Graph_algo/BFS](../Graph_algo/BFS/)
