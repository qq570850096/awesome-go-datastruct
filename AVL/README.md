# AVL Tree

## Goal

Learn strict height-balanced binary search trees and rotations.

## Prerequisites

Binary search trees and recursion.

## Core Invariant

For every node, the heights of left and right subtrees differ by at most one.

## Complexity

Search/insert/delete: Time O(log n). Space O(h) for recursion.

## Practice Tasks

- Trace LL, RR, LR, and RL rotations.
- Add deletion cases that require rebalancing.
- Compare height behavior with a plain BST.

## Test Command

```bash
go test ./AVL
```

## Related Topics

- [BinarySearch](../BinarySearch/)
- [Red-Black](../Red-Black/)
