# Generics

Reusable, type-safe containers and slice helpers using Go type parameters.

## Quick Start

```bash
go test ./BasicGo/generics
go test ./BasicGo/generics -run TestStack
```

## What You Will Learn

- How `Stack[T]` keeps one element type without casts.
- How generic functions can transform from `T` to `R`.
- How constraints such as `comparable` document required operations.
- How zero values are returned when a generic container has no item to produce.

## Concept Map

```text
type parameter T
      |
      +-- container: Stack[T]
      +-- transform: MapSlice[T, R]
      +-- predicate: FilterSlice[T]
constraint
      |
      +-- comparable when equality is needed
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Stack[T]` | LIFO container backed by a slice. | `Pop` returns zero value and `false` when empty. |
| `Push(v)` | Add one value to the top. | Amortized O(1). |
| `Pop()` | Remove the newest value. | Mutates the stack length. |
| `Len()` | Return the number of stored values. | O(1). |
| `MapSlice[T, R]` | Transform each item into another type. | Preserves input order and length. |
| `FilterSlice[T comparable]` | Keep values matching a predicate. | Preserves relative order of kept values. |

## Guided Walkthrough

1. Read `Stack[T]` as a generic version of the classic stack module.
2. Read `Pop` and notice how it creates a zero value for empty stacks.
3. Read `MapSlice` and identify why it needs two type parameters.
4. Read `FilterSlice` and ask whether `comparable` is necessary for the current implementation.

## Example

```go
var names Stack[string]
names.Push("ada")
names.Push("grace")
top, ok := names.Pop() // "grace", true
```

The compiler prevents accidentally pushing an `int` into `Stack[string]`.

## Common Pitfalls

- Reaching for `any` when a type parameter would preserve safety.
- Adding constraints that the implementation does not actually use.
- Forgetting that an empty generic container still returns a typed zero value.
- Mutating input slices in helpers that promise transformation only.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add `Peek()` to `Stack[T]` without changing length. |
| Drill | Add `ReduceSlice[T, R]` and test it with strings and integers. |
| Challenge | Compare `Stack[T]` with the non-generic `stack` module and document tradeoffs. |

## Quality Checklist

- Generic APIs express constraints precisely.
- Tests use more than one concrete type.
- Empty-container behavior is explicit.
- Helpers preserve order unless documented otherwise.

## Related Topics

- [stack](../../stack/)
- [slicemap](../slicemap/)
- [Utils/Interfaces](../../Utils/Interfaces/)
