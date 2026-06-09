# Creational Design Patterns

Creational patterns centralize object creation so callers do not have to know every construction detail, product variant, or initialization step.

## Quick Start

```bash
go test ./DesignPatterns/CreativeType
```

Run one pattern while studying it:

```bash
go test ./DesignPatterns/CreativeType -run TestGetInstance
go test ./DesignPatterns/CreativeType -run TestDirector
```

## What This Module Covers

- One-time initialization with singleton.
- Cloning existing objects with prototype.
- Simple construction dispatch with factory methods.
- Families of related products with abstract factory.
- Step-by-step assembly with builder and director roles.

## Pattern Index

| Pattern | Code | Detailed doc | Primary lesson |
| --- | --- | --- | --- |
| Singleton | [Singleton.go](Singleton.go) | [singleton.md](Docs/singleton.md) | `sync.Once` protects one shared instance. |
| Prototype | [Prototype.go](Prototype.go) | [prototype.md](Docs/prototype.md) | A new object can be cloned from an existing configured object. |
| Simple Factory | [SimpleFactory.go](SimpleFactory.go) | [simple-factory.md](Docs/simple-factory.md) | Construction choice is centralized in one factory. |
| Abstract Factory | [AbstractFactory.go](AbstractFactory.go) | [abstract-factory.md](Docs/abstract-factory.md) | A factory creates compatible products from one family. |
| Builder | [Builder.go](Builder.go) | [builder.md](Docs/builder.md) | A director coordinates construction steps for a complex product. |

## Learning Path

1. Start with [SimpleFactory.go](SimpleFactory.go) to see construction dispatch in its smallest form.
2. Read [AbstractFactory.go](AbstractFactory.go) and compare "one product choice" with "a family of product choices".
3. Read [Builder.go](Builder.go) and identify the product, builder, concrete builder, and director.
4. Read [Prototype.go](Prototype.go) and inspect what clone ownership means.
5. Finish with [Singleton.go](Singleton.go); discuss why global state is convenient but risky.

## Role Vocabulary

| Role | Meaning in this module |
| --- | --- |
| Product | Object returned to the client. |
| Factory | Object or function that decides which product to create. |
| Concrete factory | Factory for one product family. |
| Builder | Step interface for assembling a product. |
| Director | Object that calls builder steps in order. |
| Prototype | Configured object used as a cloning source. |

## Design Notes

- Use a factory when construction choice is meaningful to centralize.
- Use an abstract factory when products must remain compatible with each other.
- Use a builder when a product has required construction phases or many optional parts.
- Use prototype when copying an existing configured object is clearer than rebuilding it.
- Treat singleton as a tradeoff: it simplifies access but makes state sharing and tests harder.

## Common Pitfalls

- Adding a factory around a constructor that has no variation.
- Returning concrete types from a factory that promised an interface boundary.
- Letting a builder produce a partially initialized product.
- Implementing prototype as a shallow copy when deep copy is required.
- Hiding mutable global state behind singleton and making tests order-dependent.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add one more product to the simple factory and cover it with a test. |
| Drill | Add a second concrete builder and compare the final `Product`. |
| Challenge | Add tests proving abstract-factory products come from the same brand family. |
| Extension | Refactor one creation example to return errors for unsupported variants. |

## Quality Checklist

- Factories document supported variants.
- Builders produce complete products.
- Singleton tests do not depend on hidden mutable state from previous tests.
- New patterns are linked from [DesignPatterns README](../README.MD).

## Related Topics

- [DesignPatterns](../)
- [StructuralType](../StructuralType/)
- [BehavioralType](../BehavioralType/)
- [BasicGo/interface](../../BasicGo/interface/)
