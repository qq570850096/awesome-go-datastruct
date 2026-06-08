# SkipList

## Goal

Learn randomized layered indexing as an alternative to balanced trees.

## Prerequisites

Linked lists and probability basics.

## Core Invariant

Higher levels skip over ordered nodes while level zero contains the full sorted sequence.

## Complexity

Expected search/insert/delete: Time O(log n). Worst case: Time O(n). Space O(n).

## Practice Tasks

- Add deterministic tests around search and delete.
- Explain how random levels affect average cost.
- Compare SkipList with AVL or red-black trees.

## Test Command

```bash
go test ./skiplists
```

## Related Topics

- [AVL](../AVL/)
- [Red-Black](../Red-Black/)
