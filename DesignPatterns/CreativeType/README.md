# Creational Design Patterns

## Goal

Study patterns that control object creation, initialization, cloning, and product-family selection.

## Pattern Map

| Pattern | Code | Detailed doc | Core idea |
| --- | --- | --- | --- |
| Singleton | [Singleton.go](Singleton.go) | [singleton.md](Docs/singleton.md) | Provide one shared instance with controlled initialization. |
| Prototype | [Prototype.go](Prototype.go) | [prototype.md](Docs/prototype.md) | Create new objects by cloning an existing object. |
| Simple Factory | [SimpleFactory.go](SimpleFactory.go) | [simple-factory.md](Docs/simple-factory.md) | Centralize construction behind a factory method. |
| Abstract Factory | [AbstractFactory.go](AbstractFactory.go) | [abstract-factory.md](Docs/abstract-factory.md) | Build families of related products through one factory interface. |
| Builder | [Builder.go](Builder.go) | [builder.md](Docs/builder.md) | Construct a complex product step by step. |

## Core Invariants

- Construction code should hide only meaningful creation complexity.
- Factories should return values that satisfy the advertised interface.
- Builder steps should produce a complete and internally consistent product.

## Practice Tasks

- Add a second concrete builder and compare the resulting product.
- Add tests that prove abstract-factory products come from the same family.
- Explain where singleton state would become risky in concurrent code.

## Test Command

```bash
go test ./DesignPatterns/CreativeType
```

## Related Topics

- [DesignPatterns](../)
- [StructuralType](../StructuralType/)
- [BehavioralType](../BehavioralType/)
