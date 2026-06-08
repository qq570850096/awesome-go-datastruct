# Circular Queue Exercise

## Goal

Practice fixed-capacity queue design with wraparound indexes.

## Prerequisites

Arrays, indexes, and queue semantics.

## Core Invariant

Head and tail indexes move modulo capacity and the full/empty states remain distinguishable.

## Complexity

All core operations are Time O(1), Space O(k).

## Practice Tasks

- Test capacity one.
- Test wraparound after mixed enqueue and dequeue operations.
- Document how full and empty are represented.

## Test Command

```bash
go test ./main
```

## Related Topics

- [queue](../queue/)
