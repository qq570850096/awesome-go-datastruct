# Binary Search Tree

## Goal

Learn ordered tree storage, traversal, lookup, insertion, and deletion.

## Prerequisites

Recursion, structs, and pointers.

## Core Invariant

For every node, left-subtree values are smaller and right-subtree values are larger under this implementation.

## Complexity

Average search/insert/delete: Time O(log n). Worst case: Time O(n). Traversal: Time O(n), Space O(h).

## Practice Tasks

- Add deletion tests for leaf, one-child, and two-child nodes.
- Compare preorder, inorder, and postorder traversal output.
- Explain why sorted insertion creates a degenerate tree.

## Test Command

```bash
go test ./BinarySearch
```

## Related Topics

- [AVL](../AVL/)
- [Red-Black](../Red-Black/)
- [Segment](../Segment/)
