# Slices And Maps

Slice descriptors, backing-array sharing, in-place filtering, map counting, and a map-backed set.

## Quick Start

```bash
go test ./BasicGo/slicemap
go test ./BasicGo/slicemap -run TestShareUnderlying
```

## What You Will Learn

- How nil and empty slices differ.
- How slices share backing arrays until growth reallocates.
- How in-place filtering reuses existing storage.
- How maps support counting and set-like membership.
- How tests reveal aliasing that is otherwise easy to miss.

## Concept Map

```text
slice = pointer + length + capacity
subslice -> same backing array
append   -> may reuse or allocate
map      -> hash table reference
set      -> map[key]struct{}
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `MakeNilAndEmpty()` | Compare nil and empty slices. | Empty slice is non-nil with length 0. |
| `ShareUnderlying()` | Demonstrate backing-array sharing and append growth. | Returned slices expose aliasing behavior. |
| `FilterInPlace(nums, keep)` | Keep matching values in the input storage. | Mutates the prefix of `nums`. |
| `CountWords(text)` | Count whitespace-separated words. | Returns a map from word to count. |
| `Set` | Map-backed string set. | Membership is represented by key presence. |
| `NewSet`, `Add`, `Has`, `Remove` | Basic set operations. | Average O(1) map behavior. |

## Guided Walkthrough

1. Start with `MakeNilAndEmpty` and inspect length, capacity, and nil-ness in tests.
2. Read `ShareUnderlying`; trace which slices share storage.
3. Read `FilterInPlace`; follow the write index.
4. Read `map.go`; compare counting with membership-only storage.

## Example

```go
nums := []int{1, 2, 3, 4}
evens := FilterInPlace(nums, func(v int) bool { return v%2 == 0 })
fmt.Println(evens) // [2 4]
```

Because filtering is in-place, callers should not assume the old prefix of `nums` is unchanged.

## Common Pitfalls

- Assuming every `append` allocates a new array.
- Returning a subslice that keeps a large backing array alive.
- Mutating a slice while another variable aliases the same backing storage.
- Confusing missing map keys with zero values.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a filter test that keeps no values. |
| Drill | Add `Len()` to `Set` and test add/remove behavior. |
| Challenge | Normalize punctuation before `CountWords` and document the policy. |

## Quality Checklist

- Functions document whether they mutate input slices.
- Tests include empty, nil, duplicate, and aliasing cases.
- Map helpers define missing-key behavior.

## Related Topics

- [generics](../generics/)
- [Set](../../Set/)
- [BasicGo/basics](../basics/)
