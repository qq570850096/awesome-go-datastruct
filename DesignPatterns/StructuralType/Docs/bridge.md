# Bridge Pattern

Category: Structural

Code: [Bridge.go](../Bridge.go)

## Intent

Separate an abstraction from its implementation so both can vary independently.

## Roles

- Abstraction owns an implementation interface.
- Implementation defines the behavior to vary.
- Refined abstractions and concrete implementations evolve separately.

## Use When

- Two dimensions of variation are growing independently.
- Inheritance-style combinations would explode.

## Tradeoffs

- Adds indirection.
- Boundaries must be named clearly.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestBridge
```
