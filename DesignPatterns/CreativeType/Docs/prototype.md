# Prototype Pattern

Category: Creational

Code: [Prototype.go](../Prototype.go)

## Intent

Create new objects by cloning an existing object instead of constructing from scratch.

## Roles

- Prototype exposes clone behavior.
- Concrete prototype copies its internal state.
- Client works with the copied instance.

## Use When

- Object construction is expensive.
- The runtime object should decide how it is copied.

## Tradeoffs

- Deep copy semantics must be explicit.
- Shared references can cause subtle bugs.

## Test Command

```bash
go test ./DesignPatterns/CreativeType -run TestPrototype
```
