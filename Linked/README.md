# Singly Linked List

## Goal

Learn pointer-based list manipulation, reversal, duplicate removal, intersection checks, and common linked-list interview exercises.

## Prerequisites

Pointers, structs, and table-driven tests.

## Core Invariant

Every node owns a value and a next pointer. List operations must preserve reachability from the head and must not create unintended cycles.

## Complexity

Search: Time O(n), Space O(1). Insert or delete after a known node: Time O(1), Space O(1). Reversal: Time O(n), Space O(1).

## Practice Tasks

- Add tests for empty lists and one-node lists.
- Trace pointer updates during reversal.
- Compare iterative and recursive duplicate removal.

## Test Command

```bash
go test ./Linked
```

## Related Topics

- [DoubleLinked](../DoubleLinked/)
- [stack](../stack/)
- [queue](../queue/)
