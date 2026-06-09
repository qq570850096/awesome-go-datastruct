# Interface Utilities

## Goal

Provide a tiny comparison helper used to demonstrate how empty-interface values can be inspected and compared in a controlled way.

## Key Ideas

- `interface{}` can accept values of any dynamic type.
- A helper that accepts `interface{}` must define supported types clearly.
- Unsupported dynamic types should fail loudly or return an explicit error.

## Repository Code Map

| File | What to read for |
| --- | --- |
| Interfaces.go | `Compare` implementation for supported dynamic types. |
| Interfaces_test.go | Ordering tests and panic behavior for unsupported values. |

## Core Invariant

`Compare(a, b)` should return a stable ordering for supported same-kind values and should not silently compare unsupported inputs.

## Practice Tasks

- Add support for one additional comparable type.
- Replace panic behavior with an error-returning API and update tests.
- Compare this helper with generic constraints in `BasicGo/generics`.

## Test Command

```bash
go test ./Utils/Interfaces
```

## Related Topics

- [BasicGo/interface](../../BasicGo/interface/)
- [BasicGo/generics](../../BasicGo/generics/)
