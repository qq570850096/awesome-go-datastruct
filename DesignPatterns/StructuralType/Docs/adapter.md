# Adapter Pattern

Category: Structural

Code: [Wrapper.go](../Wrapper.go)

## Intent

Convert one interface into another interface expected by the client.

## Roles

- Target is the interface the client expects.
- Adaptee is the incompatible existing type.
- Adapter wraps the adaptee and implements the target.

## Use When

- Existing code has useful behavior but the wrong interface.
- You need compatibility without modifying the adaptee.

## Tradeoffs

- Adapters add another layer.
- Too many adapters can hide design drift.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestWrapper
```
