# Pointers

## Goal

Understand address semantics, caller-visible mutation, nil checks, swapping, and pointer-linked nodes.

## Key Ideas

- Passing a value copies it.
- Passing a pointer lets the callee mutate caller-visible state.
- Nil pointers must be checked before dereference.
- Linked structures are built by storing addresses to other nodes.

## Repository Code Map

| File | What to read for |
| --- | --- |
| pointers.go | Value increment, pointer increment, swap, linked-node construction, and traversal. |
| pointers_test.go | Tests for mutation, nil behavior, and linked-list values. |

## Core Invariant

Functions that accept pointers should define what happens for nil and should only mutate state intentionally.

## Practice Tasks

- Add a `Len` helper for linked nodes.
- Add tests for nil inputs to `Swap`.
- Draw the node addresses created by `Link(1, 2, 3)`.

## Test Command

```bash
go test ./BasicGo/pointers
```

## Related Topics

- [Linked](../../Linked/)
- [structs](../structs/)
