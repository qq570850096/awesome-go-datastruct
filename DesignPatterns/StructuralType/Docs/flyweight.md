# Flyweight Pattern

Category: Structural

Code: [Flyweight.go](../Flyweight.go)

## Intent

Share many fine-grained objects to reduce allocation and memory cost.

## Roles

- Flyweight stores intrinsic reusable state.
- Factory returns shared objects.
- Client supplies extrinsic state when using the object.

## Use When

- A large number of similar objects would otherwise be allocated.
- Most object state can be shared safely.

## Tradeoffs

- Shared state must be immutable or carefully controlled.
- Separating intrinsic and extrinsic state adds complexity.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestFlyweight
```
