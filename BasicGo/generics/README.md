# Generics

## Goal

Use type parameters to write reusable data structures and slice helpers while keeping concrete type safety.

## Key Ideas

- `Stack[T]` stores values of one chosen element type.
- Generic functions can transform from one type parameter to another.
- Constraints such as `comparable` express required operations.
- Zero values still matter when a generic container is empty.

## Repository Code Map

| File | What to read for |
| --- | --- |
| stack.go | Generic stack, map helper, and filter helper. |
| stack_test.go | Type-specific stack and slice helper tests. |

## Core Invariant

Generic helpers should avoid `any` casts when the type parameter can express the rule directly.

## Practice Tasks

- Add `Peek` to `Stack[T]`.
- Add a `ReduceSlice` helper and tests.
- Try `Stack[string]` and `Stack[int]` in the same test.

## Test Command

```bash
go test ./BasicGo/generics
```

## Related Topics

- [stack](../../stack/)
- [slicemap](../slicemap/)
