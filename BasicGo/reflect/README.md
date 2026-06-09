# Reflection

Runtime type-inspection examples using the `reflect` package.

## Quick Start

```bash
go test ./BasicGo/reflect
```

## What You Will Learn

- How `reflect.TypeOf` discovers the dynamic type of an interface value.
- How `Kind` groups concrete types into broader categories.
- Why reflection is powerful for tooling and adapters but should be used sparingly.
- How a reflection demo can be improved by returning values that tests can assert.

## Concept Map

```text
interface value -> dynamic type
reflect.TypeOf  -> Type
Type.Kind()     -> category such as Int, Float64, Slice
ordinary code   -> prefer direct types or small interfaces
```

## API Surface

| Function | Purpose | Important contract |
| --- | --- | --- |
| `CheckType(v interface{})` | Print a broad type category for numeric values. | Demonstration prints output; it does not return a value. |

## Guided Walkthrough

1. Read `CheckType` and identify the conversion from interface value to `reflect.Type`.
2. Compare `Kind()` cases for signed, unsigned, and floating-point values.
3. Read `Reflect_test.go`; it currently acts as a smoke test.
4. Consider refactoring the function to return a string, which would make tests stronger.

## Example

```go
CheckType(10)    // signed integer
CheckType(10.5)  // float
CheckType("go")  // Unknown string
```

Reflection is useful when the static type is unavailable. If a small interface can describe the behavior, prefer that interface.

## Common Pitfalls

- Using reflection to avoid designing a clear interface.
- Forgetting that `reflect.TypeOf(nil)` returns nil.
- Assuming `Kind` and concrete type name are the same thing.
- Writing tests that only execute reflection code without asserting output.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add nil handling to `CheckType`. |
| Drill | Return the category string instead of printing it. |
| Challenge | Add slice and map classification and compare it with `TypeName` in `basics`. |

## Quality Checklist

- Reflection examples are isolated and intentional.
- Tests assert returned behavior when possible.
- New cases distinguish concrete type from kind.

## Related Topics

- [interface](../interface/)
- [structs](../structs/)
- [lowlevel](../lowlevel/)
