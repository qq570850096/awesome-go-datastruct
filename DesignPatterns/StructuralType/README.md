# Structural Design Patterns

## Goal

Study patterns that compose objects, adapt interfaces, share heavy state, protect access, and simplify subsystems.

## Pattern Map

| Pattern | Code | Detailed doc | Core idea |
| --- | --- | --- | --- |
| Adapter | [Wrapper.go](Wrapper.go) | [adapter.md](Docs/adapter.md) | Convert one interface into another expected interface. |
| Bridge | [Bridge.go](Bridge.go) | [bridge.md](Docs/bridge.md) | Separate abstraction from implementation so both can vary. |
| Composite | [Component.go](Component.go) | [composite.md](Docs/composite.md) | Treat leaves and groups through one component interface. |
| Decorator | [Decorator.go](Decorator.go) | [decorator.md](Docs/decorator.md) | Add behavior by wrapping an object. |
| Facade | [Facade.go](Facade.go) | [facade.md](Docs/facade.md) | Provide a simple entrypoint over a subsystem. |
| Filter | [Filter.go](Filter.go) | [filter.md](Docs/filter.md) | Encapsulate reusable selection criteria. |
| Flyweight | [Flyweight.go](Flyweight.go) | [flyweight.md](Docs/flyweight.md) | Share intrinsic state across many objects. |
| Proxy | [Proxy.go](Proxy.go) | [proxy.md](Docs/proxy.md) | Control access to a real subject. |

## Core Invariants

- Wrappers should preserve the contract of the object they expose.
- Composite operations should work uniformly for leaves and groups where the interface promises it.
- Shared flyweight state should not accidentally store caller-specific data.

## Practice Tasks

- Add a test proving an adapter can be used through the target interface.
- Extend the facade with one more subsystem operation.
- Compare decorator and proxy: both wrap, but their intent differs.

## Test Command

```bash
go test ./DesignPatterns/StructuralType
```

## Related Topics

- [DesignPatterns](../)
- [CreativeType](../CreativeType/)
- [BehavioralType](../BehavioralType/)
