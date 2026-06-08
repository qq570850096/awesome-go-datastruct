# Iterator Pattern

Category: Behavioral

Code: [Iterator.go](../Iterator.go)

## Intent

Provide a uniform way to traverse a collection without exposing its internal storage.

## Roles

- Iterator exposes traversal operations.
- Aggregate creates iterators.
- Concrete iterator tracks traversal state.

## Use When

- Clients need traversal without depending on collection structure.
- Different collections should share one traversal style.

## Tradeoffs

- Simple Go slices often need no custom iterator.
- Iterator state must remain valid if the collection mutates.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestIterator
```
