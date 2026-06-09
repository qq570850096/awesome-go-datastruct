# Slices And Maps

## Goal

Understand slice descriptors, backing-array sharing, in-place filtering, map counting, and map-backed sets.

## Key Ideas

- A nil slice and an empty slice behave similarly in many operations but are not the same value.
- Slices can share a backing array until append growth allocates a new one.
- In-place filtering reuses storage and writes kept values into the original slice prefix.
- A map with `struct{}` values is a common lightweight set representation.

## Repository Code Map

| File | What to read for |
| --- | --- |
| slice.go | Nil versus empty slices, backing-array sharing, and in-place filtering. |
| map.go | Word counts and a map-backed string set. |
| slicemap_test.go | Tests for slice and map behavior. |

## Core Invariant

When returning a slice, callers should know whether it may share storage with the input.

## Practice Tasks

- Add a filter test that keeps no values.
- Add a word-count case with repeated punctuation after normalizing input.
- Add a `Len` method to `Set`.

## Test Command

```bash
go test ./BasicGo/slicemap
```

## Related Topics

- [generics](../generics/)
- [Set](../../Set/)
