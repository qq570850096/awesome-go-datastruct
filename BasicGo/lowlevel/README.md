# Low-Level Layout

## Goal

Inspect struct size, alignment, and field offsets with `unsafe` so memory layout becomes observable.

## Key Ideas

- A struct may contain padding between fields.
- Alignment requirements influence total struct size.
- `unsafe.Sizeof`, `unsafe.Alignof`, and `unsafe.Offsetof` are inspection tools.
- Low-level knowledge helps explain performance, but it should not leak into ordinary code unnecessarily.

## Repository Code Map

| File | What to read for |
| --- | --- |
| lowlevel.go | `Layout`, `LayoutInfo`, and layout inspection. |
| lowlevel_test.go | Assertions that make size and offset behavior visible. |

## Core Invariant

Layout observations should be treated as implementation-specific details unless the code explicitly owns that low-level contract.

## Practice Tasks

- Reorder fields in a new struct and compare size.
- Add alignment checks for another primitive type.
- Explain why padding can change memory usage for large slices of structs.

## Test Command

```bash
go test ./BasicGo/lowlevel
```

## Related Topics

- [structs](../structs/)
- [pointers](../pointers/)
