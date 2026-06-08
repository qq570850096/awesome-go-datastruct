# Segment Tree

## Goal

Learn range aggregation and update through a tree over array intervals.

## Prerequisites

Arrays, recursion, and merge functions.

## Core Invariant

Each internal node stores the merged value of its left and right child intervals.

## Complexity

Build: Time O(n). Query: Time O(log n). Update: Time O(log n). Space O(n).

## Practice Tasks

- Add left-only and right-only query tests.
- Test updates at the first and last index.
- Explain how the merger controls behavior.

## Test Command

```bash
go test ./Segment
```

## Related Topics

- [BinarySearch](../BinarySearch/)
- [Sorts](../Sorts/)
