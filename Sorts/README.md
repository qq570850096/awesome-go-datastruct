# Sorting Algorithms

## Goal

Compare common sorting algorithms through Go implementations and tests.

## Prerequisites

Arrays, loops, recursion, and comparison functions.

## Core Invariant

Output must be ordered and contain exactly the same multiset of elements as input.

## Complexity

Bubble/insertion/selection: Time O(n^2). Merge/quick average: Time O(n log n). Bucket/counting depend on value range. Space varies by algorithm.

## Practice Tasks

- Add tests for duplicates and already-sorted input.
- Add benchmark cases for random and sorted input.
- Document stability for each algorithm.

## Test Command

```bash
go test ./Sorts
```

## Related Topics

- [Heap](../Heap/)
- [Graph_algo](../Graph_algo/)
