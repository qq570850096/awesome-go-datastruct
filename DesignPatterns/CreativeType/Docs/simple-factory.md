# Simple Factory Pattern

Category: Creational

Code: [SimpleFactory.go](../SimpleFactory.go)

## Intent

Centralize object creation behind a small factory function.

## Roles

- Factory chooses the concrete type.
- Product interface or base type hides concrete details.
- Client asks the factory for a product.

## Use When

- Creation logic is simple but repeated.
- Callers should not know concrete constructors.

## Tradeoffs

- The factory can become a switch-heavy hotspot.
- Adding products may require modifying the factory.

## Test Command

```bash
go test ./DesignPatterns/CreativeType -run TestSimpleFactory
```
