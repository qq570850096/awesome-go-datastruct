# Interface Utilities

A small comparison helper used to discuss dynamic values, supported type policies, and the tradeoff between `interface{}` and generics.

## Quick Start

```bash
go test ./Utils/Interfaces
go test ./Utils/Interfaces -run TestCompare
```

## What This Module Covers

- Accepting values through `interface{}`.
- Comparing supported dynamic types through a single helper.
- Defining behavior for unsupported inputs.
- Testing panic behavior intentionally.
- Comparing this older style with generic constraints from `BasicGo/generics`.

## API Surface

| Function | Purpose | Important contract |
| --- | --- | --- |
| `Compare(a, b interface{}) int` | Compare two supported values and return ordering. | Returns a negative, zero, or positive value; panics for unsupported inputs. |

## Behavioral Contract

```text
Compare(a, b) < 0  -> a sorts before b
Compare(a, b) == 0 -> a and b are equal by the helper policy
Compare(a, b) > 0  -> a sorts after b
unsupported input  -> panic
```

The tests are part of the contract. Before changing panic behavior or adding supported types, update tests first so callers can see the new policy.

## Reading Guide

1. Read `Interfaces.go` and identify every supported dynamic type.
2. Read `Interfaces_test.go` and map each test case to one return shape.
3. Compare this helper with a possible generic function such as `Compare[T constraints.Ordered]`.
4. Decide whether a future API should keep panics or return `(int, error)`.

## Common Pitfalls

- Silently comparing unsupported types by their formatted string output.
- Returning inconsistent signs for equal or ordered values.
- Mixing type policies in tests and implementation.
- Using `interface{}` when a generic constraint would make unsupported values impossible at compile time.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add support for one more numeric type and cover it with tests. |
| Drill | Replace panic behavior with `(int, error)` and document the migration. |
| Challenge | Implement a generic ordered comparator in a separate file and compare APIs. |

## Quality Checklist

- Supported types are documented in tests.
- Unsupported behavior is explicit.
- Return sign conventions match sorting expectations.
- New helpers remain small enough not to hide the concept being taught.

## Related Topics

- [BasicGo/interface](../../BasicGo/interface/)
- [BasicGo/generics](../../BasicGo/generics/)
- [Sorts](../../Sorts/)
