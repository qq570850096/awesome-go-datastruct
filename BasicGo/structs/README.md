# Structs And Methods

Struct fields, value receivers, pointer receivers, embedding, promoted methods, and JSON tags.

## Quick Start

```bash
go test ./BasicGo/structs
go test ./BasicGo/structs -run TestValueVsPointerReceiver
```

## What You Will Learn

- How structs group named fields into a concrete model.
- How value receivers operate on a copy.
- How pointer receivers mutate caller-visible state.
- How embedding exposes promoted fields and methods.
- How JSON tags map Go field names to encoded names.

## Concept Map

```text
struct field     -> named data
value receiver   -> copy
pointer receiver -> mutate original
embedding        -> compose behavior
tag              -> metadata for libraries
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `User` | Demonstrates receiver mutation rules. | `RenameValue` does not mutate caller state; `RenamePointer` does. |
| `ChangeUserValue` | Mutate a copy. | Caller does not observe the change. |
| `ChangeUserPointer` | Mutate through pointer. | Caller observes the change. |
| `Logger` | Embedded behavior provider. | `Log` formats a message. |
| `Service` | Embeds `Logger`. | Uses promoted logging behavior. |
| `Account` | JSON tag example. | Encodes with tagged field names. |
| `EncodeAccount`, `DecodeAccount` | JSON round-trip helpers. | Return errors from `encoding/json`. |

## Guided Walkthrough

1. Read `receiver.go` and compare value and pointer receivers.
2. Read `embedding.go`; identify where `Service` gains logging behavior.
3. Read `tag.go`; map Go field names to JSON keys.
4. Read tests and notice how each section protects a different struct concept.

## Example

```go
u := User{Name: "old"}
u.RenamePointer("new")
fmt.Println(u.Name) // new
```

The pointer receiver lets the method modify the original `User`.

## Common Pitfalls

- Expecting a value receiver to mutate caller-visible state.
- Mixing value and pointer receivers without a reason.
- Forgetting that only exported fields are encoded by `encoding/json`.
- Overusing embedding when an explicit field would be clearer.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add an exported and an unexported field to `Account` and compare JSON output. |
| Drill | Add a read-only method with a value receiver and explain why it is safe. |
| Challenge | Build a small `Repository` struct that embeds a logger and exposes one testable method. |

## Quality Checklist

- Receiver choice matches mutation and copying cost.
- JSON tests assert field names, not only successful encoding.
- Embedded behavior is used to clarify composition.

## Related Topics

- [pointers](../pointers/)
- [interface](../interface/)
- [errors](../errors/)
