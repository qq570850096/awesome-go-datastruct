# Low-Level Layout

Memory-layout inspection examples using `unsafe.Sizeof`, `unsafe.Alignof`, and `unsafe.Offsetof`.

## Quick Start

```bash
go test ./BasicGo/lowlevel
```

## What You Will Learn

- How struct field order influences padding.
- How alignment requirements affect total size.
- How to inspect field offsets safely in a small teaching example.
- Why layout knowledge is useful for performance but should not dominate ordinary code.

## Concept Map

```text
field type -> alignment requirement
field order -> padding
struct size -> data bytes + padding bytes
offset      -> byte position from struct start
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `Layout` | Example struct with `bool`, `int32`, and `string`. | Field order is intentionally visible. |
| `LayoutInfo` | Returned size, alignment, and selected offsets. | Uses `uintptr` values from `unsafe`. |
| `InspectLayout()` | Collect layout details for `Layout`. | Observes the current compiler/platform layout. |

## Guided Walkthrough

1. Read the field order in `Layout`.
2. Use `InspectLayout` to see total size and alignment.
3. Compare `CountOffset` and `NameOffset` with the field order.
4. Read the test and identify which assertions depend on Go's current layout rules.

## Example

```go
info := InspectLayout()
fmt.Println(info.Size, info.Align, info.CountOffset, info.NameOffset)
```

Use this as an observation tool, not as a reason to introduce `unsafe` into ordinary modules.

## Common Pitfalls

- Treating observed layout as portable across every architecture without checking.
- Reordering fields only for size while making business code harder to read.
- Using `unsafe` when `reflect` or ordinary field access would do.
- Forgetting that a slice of structs multiplies padding cost by element count.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add another struct with reordered fields and compare sizes. |
| Drill | Inspect offsets for `bool`, `int64`, and `byte` combinations. |
| Challenge | Estimate memory usage for one million values of two layouts. |

## Quality Checklist

- Low-level examples are isolated from high-level modules.
- Tests document platform-sensitive assumptions.
- Comments explain why `unsafe` is being used.

## Related Topics

- [structs](../structs/)
- [pointers](../pointers/)
- [reflect](../reflect/)
