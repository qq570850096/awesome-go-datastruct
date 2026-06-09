# Pointers

Address semantics, caller-visible mutation, nil handling, swapping, and pointer-linked nodes.

## Quick Start

```bash
go test ./BasicGo/pointers
go test ./BasicGo/pointers -run TestLinkedNodes
```

## What You Will Learn

- How passing a value differs from passing its address.
- How pointer parameters allow caller-visible mutation.
- How nil checks define safe pointer APIs.
- How linked structures are built by storing addresses to other nodes.

## Concept Map

```text
value assignment -> copy
&value           -> address
*pointer         -> value at address
nil pointer      -> no address
node.Next        -> link to another node
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Node` | Minimal singly linked node. | `Next` may be nil at the tail. |
| `IncrementValue(n)` | Increment a copy and return it. | Does not mutate caller state. |
| `IncrementPointer(n)` | Increment through `*int`. | Returns `false` for nil. |
| `Swap(a, b)` | Swap two pointed-to integers. | Returns `false` when either pointer is nil. |
| `Link(values...)` | Build a linked node chain. | Returns nil for no values. |
| `Values(head)` | Traverse a chain into a slice. | Stops at nil. |

## Guided Walkthrough

1. Compare `IncrementValue` and `IncrementPointer`.
2. Read `Swap` and identify the nil-protection boundary.
3. Read `Link`; follow how the tail pointer moves.
4. Read `Values`; it is the simplest traversal pattern for linked structures.

## Example

```go
a, b := 1, 2
ok := Swap(&a, &b)
fmt.Println(ok, a, b) // true 2 1
```

The function receives addresses, so it can mutate the caller's variables.

## Common Pitfalls

- Expecting a value parameter to mutate the caller's variable.
- Dereferencing nil before checking it.
- Losing the head pointer while building a linked list.
- Returning internal pointers from APIs without documenting ownership.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add `Len(head *Node) int`. |
| Drill | Add tests for nil arguments to `IncrementPointer` and `Swap`. |
| Challenge | Implement `Reverse(head *Node) *Node` and compare it with the `Linked` module. |

## Quality Checklist

- Pointer APIs specify nil behavior.
- Tests check both mutation and non-mutation.
- Linked traversal terminates at nil.
- Examples make aliasing visible.

## Related Topics

- [Linked](../../Linked/)
- [structs](../structs/)
- [lowlevel](../lowlevel/)
