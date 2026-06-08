# Heap And Priority Queue

## Goal

Learn array-backed heap operations and priority-queue behavior.

## Prerequisites

Arrays, tree indexes, and comparison functions.

## Core Invariant

Each parent has priority at least as high as its children for a max-heap.

## Complexity

Add: Time O(log n). Remove max: Time O(log n). Peek max: Time O(1). Heapify: Time O(n).

## Practice Tasks

- Add repeated insert/remove tests.
- Add a benchmark for heap construction.
- Explain parent, left-child, and right-child index formulas.

## Test Command

```bash
go test ./Heap
```

## Related Topics

- [Sorts](../Sorts/)
- [Graph_algo](../Graph_algo/)
