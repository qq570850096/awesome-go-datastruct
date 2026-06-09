# Interfaces

## Goal

Learn how Go interfaces describe behavior, how `any` carries dynamic values, and how errors participate in interface-based design.

## Key Ideas

- A type satisfies an interface implicitly by implementing its methods.
- Interfaces should usually be small and behavior-oriented.
- `any` is useful at boundaries, but concrete behavior should be recovered carefully.
- Error values are ordinary interface values.

## Repository Code Map

| File | What to read for |
| --- | --- |
| shape.go | `Shape`, rectangle and circle implementations, and polymorphic area calculation. |
| empty.go | `any` and type-switch descriptions. |
| error_interface.go | Custom error type and wrapped temporary error. |
| interface_test.go | Interface dispatch, `any`, and error wrapping tests. |

## Core Invariant

Callers should depend on the smallest behavior they need, not on concrete implementation details.

## Practice Tasks

- Add a triangle shape and include it in `TotalArea`.
- Extend `Describe` with one additional dynamic type.
- Add an `errors.As` assertion for `OpError`.

## Test Command

```bash
go test ./BasicGo/interface
```

## Related Topics

- [structs](../structs/)
- [errors](../errors/)
