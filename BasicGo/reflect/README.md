# Reflection

## Goal

Introduce `reflect` as a way to inspect dynamic type information when static interfaces are not enough.

## Key Ideas

- Reflection inspects runtime type and value descriptors.
- It is useful for tooling, serialization, and generic adapters.
- Reflection is less direct than ordinary interfaces and should be reserved for cases that need it.
- Tests should assert observable behavior rather than only printing reflection output.

## Repository Code Map

| File | What to read for |
| --- | --- |
| Reflect.go | Basic runtime type inspection. |
| Reflect_test.go | Smoke test for the reflection example. |

## Core Invariant

Prefer ordinary types and interfaces first; use reflection only when the program must inspect values whose static shape is not known.

## Practice Tasks

- Return a type description string instead of only printing it.
- Add cases for pointer and slice values.
- Compare `reflect.TypeOf` with a type switch from the interface module.

## Test Command

```bash
go test ./BasicGo/reflect
```

## Related Topics

- [interface](../interface/)
- [structs](../structs/)
