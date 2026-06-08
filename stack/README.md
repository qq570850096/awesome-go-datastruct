# Stack

## Goal

Learn last-in first-out behavior with a compact Go implementation.

## Prerequisites

Slices and basic tests.

## Core Invariant

The most recently pushed item is the next item popped.

## Complexity

Push: Time O(1) amortized. Pop: Time O(1). Space O(n).

## Practice Tasks

- Add an empty-pop test.
- Add a multi-item order test.
- Compare stack behavior with DFS recursion.

## Test Command

```bash
go test ./stack
```

## Related Topics

- [queue](../queue/)
- [Graph_algo/DFS](../Graph_algo/DFS/)
