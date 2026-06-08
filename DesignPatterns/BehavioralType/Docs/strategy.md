# Strategy Pattern

Category: Behavioral

Code: [Strategy.go](../Strategy.go)

## Intent

Define a family of algorithms and make them interchangeable behind one interface.

## Roles

- Strategy defines the algorithm interface.
- Concrete strategies implement alternatives.
- Context delegates behavior to the selected strategy.

## Use When

- The same operation has multiple algorithms.
- Callers should switch behavior without branching everywhere.

## Tradeoffs

- Many strategies can make configuration complex.
- The strategy interface should stay focused.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestStrategy
```
