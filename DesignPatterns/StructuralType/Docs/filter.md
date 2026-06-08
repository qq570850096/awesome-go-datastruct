# Filter Pattern

Category: Structural

Code: [Filter.go](../Filter.go)

## Intent

Encapsulate selection criteria so filters can be reused and composed.

## Roles

- Criteria defines a filter contract.
- Concrete criteria selects matching items.
- Client combines criteria as needed.

## Use When

- Selection rules are repeated.
- Rules need to be composed or swapped.

## Tradeoffs

- Small filters are easy to over-abstract.
- Composed filters should remain readable.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestFilter
```
