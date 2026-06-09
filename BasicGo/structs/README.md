# Structs And Methods

## Goal

Learn struct fields, value receivers, pointer receivers, embedding, promoted methods, and JSON tags.

## Key Ideas

- Structs group related data under named fields.
- Value receivers operate on a copy.
- Pointer receivers can mutate the original value.
- Embedding reuses behavior through promoted fields and methods.
- Tags provide metadata for packages such as `encoding/json`.

## Repository Code Map

| File | What to read for |
| --- | --- |
| receiver.go | Value versus pointer receiver behavior. |
| embedding.go | Embedded logger and service composition. |
| tag.go | JSON encoding and decoding with field tags. |
| structs_test.go | Tests for receiver, embedding, and tag behavior. |

## Core Invariant

Choose pointer receivers when the method mutates state or copying the value would be misleading or expensive.

## Practice Tasks

- Add an unexported field and observe JSON encoding behavior.
- Add a method that reads state but does not mutate it.
- Compare `ChangeUserValue` and `ChangeUserPointer` with a diagram.

## Test Command

```bash
go test ./BasicGo/structs
```

## Related Topics

- [pointers](../pointers/)
- [interface](../interface/)
