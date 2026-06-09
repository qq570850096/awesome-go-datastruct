# Interfaces

Behavior-oriented examples for implicit interface satisfaction, `any`, custom errors, and wrapped error values.

## Quick Start

```bash
go test ./BasicGo/interface
go test ./BasicGo/interface -run TestShapeInterface
```

## What You Will Learn

- How concrete types satisfy interfaces without explicit declarations.
- Why small interfaces are easier to reuse and test.
- How `any` stores a dynamic type and value.
- How error values are ordinary interface values.
- How wrapping preserves an inspectable error cause.

## Concept Map

```text
concrete type + required methods -> satisfies interface
interface value                  -> dynamic type + dynamic value
error                            -> interface with Error() string
any                              -> no behavior promised
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Shape` | Area behavior contract. | Implemented by any type with `Area() float64`. |
| `Rect`, `Circle` | Concrete shapes. | Value receivers make them easy to pass by value. |
| `TotalArea(shapes)` | Polymorphic area aggregation. | Depends only on `Shape`. |
| `Describe(v any)` | Type-switch demonstration. | Returns descriptions for supported dynamic types. |
| `OpError` | Custom error carrying operation detail. | Implements `error`. |
| `ErrTemporary` | Sentinel error category. | Intended for `errors.Is`. |
| `WrapAsTemporary(op)` | Wrap custom error around temporary sentinel. | Preserves inspectability. |

## Guided Walkthrough

1. Start with `shape.go`; it is the cleanest example of interface-driven design.
2. Move to `empty.go` and compare `any` with `Shape`.
3. Finish with `error_interface.go`; errors are interfaces too.
4. Read tests and notice how they assert behavior rather than concrete types where possible.

## Example

```go
shapes := []Shape{Rect{Width: 2, Height: 3}, Circle{Radius: 1}}
fmt.Println(TotalArea(shapes))
```

`TotalArea` does not know or care which concrete shapes it receives.

## Common Pitfalls

- Designing interfaces with too many methods too early.
- Accepting `any` when a small interface communicates the actual requirement.
- Type asserting without checking the `ok` result.
- Wrapping errors but then testing them with string comparisons.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add `Triangle` and include it in `TotalArea` tests. |
| Drill | Extend `Describe` with a slice or map case. |
| Challenge | Add a `Perimeter` interface and decide whether it belongs with `Shape`. |

## Quality Checklist

- Interfaces describe behavior, not storage.
- Callers depend on the narrowest useful contract.
- Error categories are testable with `errors.Is` or `errors.As`.
- Tests include at least two concrete implementations.

## Related Topics

- [structs](../structs/)
- [errors](../errors/)
- [DesignPatterns](../../DesignPatterns/)
