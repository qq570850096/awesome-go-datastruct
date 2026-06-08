# Decorator Pattern

Category: Structural

Code: [Decorator.go](../Decorator.go)

## Intent

Add behavior to an object dynamically while preserving the same interface.

## Roles

- Component defines the interface.
- Concrete component provides base behavior.
- Decorator wraps a component and adds behavior.

## Use When

- Behavior should be composed at runtime.
- Subclassing or branching would be too rigid.

## Tradeoffs

- Many wrappers can make call flow harder to trace.
- Tests should assert wrapper order.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestDecorator
```
