# Doubly Linked List And Caches

## Goal

Practice bidirectional pointer updates and cache eviction policies through LRU, LFU, and FIFO examples.

## Prerequisites

Singly linked lists, maps, and pointer receivers.

## Core Invariant

For every node, the next node must point back through prev and the previous node must point forward through next. Cache indexes and list order must agree.

## Complexity

LRU Get/Put: Time O(1), Space O(capacity). LFU Get/Put: average Time O(1) in this model. FIFO Get/Put: Time O(1).

## Practice Tasks

- Add tests for capacity one.
- Assert eviction order after repeated hits.
- Check that deleted nodes are no longer reachable.

## Test Command

```bash
go test ./DoubleLinked
```

## Related Topics

- [Linked](../Linked/)
- [Heap](../Heap/)
