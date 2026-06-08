# Builder Pattern

Category: Creational

Code: [Builder.go](../Builder.go)

## Intent

Separate stepwise construction from the final object representation.

## Roles

- Builder defines construction steps.
- Concrete builder stores intermediate state.
- Director controls the build sequence.
- Product is the constructed result.

## Use When

- Objects have many optional parts.
- Construction order matters.

## Tradeoffs

- Can be verbose for simple objects.
- Builder state must be kept consistent.

## Test Command

```bash
go test ./DesignPatterns/CreativeType -run TestBuilder
```
