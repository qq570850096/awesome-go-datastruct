# Structural Design Patterns

Structural patterns compose objects into larger shapes while keeping interfaces usable, responsibilities separated, and heavy state shared when appropriate.

## Quick Start

```bash
go test ./DesignPatterns/StructuralType
```

Run one pattern while studying it:

```bash
go test ./DesignPatterns/StructuralType -run TestAdapter
go test ./DesignPatterns/StructuralType -run TestComposite
```

## What This Module Covers

- Interface conversion with adapter.
- Independent abstraction and implementation hierarchies with bridge.
- Tree-shaped object structures with composite.
- Dynamic behavior extension with decorator.
- Simplified subsystem entrypoints with facade.
- Reusable selection criteria with filter.
- Intrinsic-state sharing with flyweight.
- Controlled access with proxy.

## Pattern Index

| Pattern | Code | Detailed doc | Primary lesson |
| --- | --- | --- | --- |
| Adapter | [Wrapper.go](Wrapper.go) | [adapter.md](Docs/adapter.md) | Wrap an incompatible object so it satisfies the target interface. |
| Bridge | [Bridge.go](Bridge.go) | [bridge.md](Docs/bridge.md) | Split abstraction from implementation so both vary independently. |
| Composite | [Component.go](Component.go) | [composite.md](Docs/composite.md) | Treat leaves and groups through one component interface. |
| Decorator | [Decorator.go](Decorator.go) | [decorator.md](Docs/decorator.md) | Add behavior by wrapping a compatible object. |
| Facade | [Facade.go](Facade.go) | [facade.md](Docs/facade.md) | Hide subsystem complexity behind a small entrypoint. |
| Filter | [Filter.go](Filter.go) | [filter.md](Docs/filter.md) | Package selection logic as reusable criteria objects. |
| Flyweight | [Flyweight.go](Flyweight.go) | [flyweight.md](Docs/flyweight.md) | Share intrinsic state and keep extrinsic state outside. |
| Proxy | [Proxy.go](Proxy.go) | [proxy.md](Docs/proxy.md) | Control access before delegating to a real subject. |

## Learning Path

1. Start with [Wrapper.go](Wrapper.go). Adapter is the most concrete example of "make this fit that interface".
2. Read [Decorator.go](Decorator.go) and [Proxy.go](Proxy.go) together; both wrap, but their intent differs.
3. Read [Facade.go](Facade.go) and compare "simplify a subsystem" with "wrap one object".
4. Read [Component.go](Component.go) for tree-style composition.
5. Finish with [Bridge.go](Bridge.go), [Filter.go](Filter.go), and [Flyweight.go](Flyweight.go).

## Role Vocabulary

| Role | Meaning in this module |
| --- | --- |
| Target interface | Interface the client wants to use. |
| Adapter | Object that translates target calls to adaptee calls. |
| Component | Shared interface for leaf and composite nodes. |
| Decorator | Wrapper that adds behavior while preserving interface compatibility. |
| Facade | Coarse-grained entrypoint over a subsystem. |
| Proxy | Surrogate that controls access to a real subject. |
| Flyweight | Shared object containing intrinsic state. |

## Design Notes

- Adapter is about compatibility; decorator is about adding behavior; proxy is about access control.
- Composite works best when clients can treat single items and groups uniformly.
- Facade should simplify common use cases without preventing direct subsystem use when needed.
- Flyweight examples must keep user-specific state outside the shared object.

## Common Pitfalls

- Calling every wrapper a decorator even when it controls access or translates interfaces.
- Putting extrinsic, caller-specific state into a flyweight cache.
- Giving leaf objects composite-only behavior that silently does nothing without documentation.
- Letting a facade become a dumping ground for unrelated operations.
- Testing only construction instead of delegated behavior.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a test proving `Adapter` can be used through the `Duck` interface. |
| Drill | Add one more subsystem method to `Facade` and assert the call path. |
| Challenge | Add removal or traversal assertions to the composite example. |
| Extension | Implement a small cache-backed proxy and document what it controls. |

## Quality Checklist

- Wrappers preserve or intentionally translate the expected contract.
- Tests assert delegation and state changes.
- Pattern intent is visible from the code shape.
- New patterns are linked from [DesignPatterns README](../README.MD).

## Related Topics

- [DesignPatterns](../)
- [CreativeType](../CreativeType/)
- [BehavioralType](../BehavioralType/)
- [BasicGo/interface](../../BasicGo/interface/)
