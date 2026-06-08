# Abstract Factory Pattern

Category: Creational

Code: [AbstractFactory.go](../AbstractFactory.go)

## Intent

Create families of related products without binding callers to concrete product classes.

## Roles

- Abstract factory defines creation methods.
- Concrete factory creates a product family.
- Abstract products define common behavior.
- Client depends on the factory and product abstractions.

## Use When

- Products must be used as compatible families.
- The system should switch product families as a unit.

## Tradeoffs

- More interfaces and types are required.
- Adding a new product kind affects all factories.

## Test Command

```bash
go test ./DesignPatterns/CreativeType -run TestAbsFactory
```
