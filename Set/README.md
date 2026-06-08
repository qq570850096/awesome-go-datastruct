# Set

## Goal

Practice uniqueness semantics with list-backed and tree-backed implementations.

## Prerequisites

Lists, trees, and equality checks.

## Core Invariant

A set contains no duplicate logical values.

## Complexity

List-backed operations are usually Time O(n). Tree-backed operations are average Time O(log n), worst Time O(n) without balancing.

## Practice Tasks

- Add duplicate insertion tests.
- Compare list and tree implementations.
- Document ordering expectations if any.

## Test Command

```bash
go test ./Set
```

## Related Topics

- [Linked](../Linked/)
- [BinarySearch](../BinarySearch/)
