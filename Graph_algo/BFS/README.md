# Breadth-First Search

## Goal

Learn level-order graph traversal using a queue.

## Prerequisites

Queue behavior and graph representation.

## Core Invariant

A vertex is enqueued only after it is marked visited, preventing duplicate processing.

## Complexity

Time O(V + E), Space O(V).

## Practice Tasks

- Trace the queue for a small graph.
- Add disconnected graph tests.
- Compare BFS distance behavior with DFS traversal order.

## Test Command

```bash
go test ./Graph_algo/BFS
```

## Related Topics

- [queue](../../queue/)
- [Graph_algo/search](../search/)
