# Red-Black Tree

## Goal

Learn color-balanced search trees and rotation/color-flip maintenance.

## Prerequisites

Binary search trees and tree rotations.

## Core Invariant

The root is black, red links do not appear consecutively, and every root-to-leaf path has the same black height.

## Complexity

Search/insert/delete: Time O(log n). Space O(h) for recursion.

## Practice Tasks

- Add tests for color flips.
- Compare insertion behavior with AVL.
- Document which red-black variant the implementation follows.

## Test Command

```bash
go test ./Red-Black
```

## Related Topics

- [BinarySearch](../BinarySearch/)
- [AVL](../AVL/)
