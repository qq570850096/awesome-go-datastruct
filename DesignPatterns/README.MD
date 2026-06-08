# Design Patterns In Go

This module implements classic GoF design patterns in Go. The goal is to learn responsibility boundaries and object collaboration, not to memorize pattern names.

## Pattern Groups

| Group | Patterns | Command |
| --- | --- | --- |
| Creational | Singleton, Prototype, Simple Factory, Abstract Factory, Builder | `go test ./DesignPatterns/CreativeType` |
| Structural | Flyweight, Proxy, Decorator, Facade, Bridge, Composite, Filter, Adapter | `go test ./DesignPatterns/StructuralType` |
| Behavioral | Mediator, Command, Memento, Template Method, State, Strategy, Observer, Interpreter, Chain of Responsibility, Iterator | `go test ./DesignPatterns/BehavioralType` |
| Compound | A simulation that combines multiple patterns | `go test ./DesignPatterns` |

## Creational Patterns

| Pattern | Code | Doc | Typical use |
| --- | --- | --- | --- |
| Singleton | [Singleton.go](CreativeType/Singleton.go) | [singleton.md](CreativeType/Docs/singleton.md) | One shared instance |
| Prototype | [Prototype.go](CreativeType/Prototype.go) | [prototype.md](CreativeType/Docs/prototype.md) | Clone existing objects |
| Simple Factory | [SimpleFactory.go](CreativeType/SimpleFactory.go) | [simple-factory.md](CreativeType/Docs/simple-factory.md) | Centralized construction |
| Abstract Factory | [AbstractFactory.go](CreativeType/AbstractFactory.go) | [abstract-factory.md](CreativeType/Docs/abstract-factory.md) | Families of related products |
| Builder | [Builder.go](CreativeType/Builder.go) | [builder.md](CreativeType/Docs/builder.md) | Stepwise construction |

## Structural Patterns

| Pattern | Code | Doc | Typical use |
| --- | --- | --- | --- |
| Flyweight | [Flyweight.go](StructuralType/Flyweight.go) | [flyweight.md](StructuralType/Docs/flyweight.md) | Share many small objects |
| Proxy | [Proxy.go](StructuralType/Proxy.go) | [proxy.md](StructuralType/Docs/proxy.md) | Control access |
| Decorator | [Decorator.go](StructuralType/Decorator.go) | [decorator.md](StructuralType/Docs/decorator.md) | Add behavior dynamically |
| Facade | [Facade.go](StructuralType/Facade.go) | [facade.md](StructuralType/Docs/facade.md) | Simplify subsystem access |
| Bridge | [Bridge.go](StructuralType/Bridge.go) | [bridge.md](StructuralType/Docs/bridge.md) | Separate abstraction from implementation |
| Composite | [Component.go](StructuralType/Component.go) | [composite.md](StructuralType/Docs/composite.md) | Tree-like object structures |
| Filter | [Filter.go](StructuralType/Filter.go) | [filter.md](StructuralType/Docs/filter.md) | Reusable selection criteria |
| Adapter | [Wrapper.go](StructuralType/Wrapper.go) | [adapter.md](StructuralType/Docs/adapter.md) | Convert one interface into another |

## Behavioral Patterns

| Pattern | Code | Doc | Typical use |
| --- | --- | --- | --- |
| Mediator | [Mediator.go](BehavioralType/Mediator.go) | [mediator.md](BehavioralType/Docs/mediator.md) | Coordinate objects through a hub |
| Command | [Command.go](BehavioralType/Command.go) | [command.md](BehavioralType/Docs/command.md) | Encapsulate requests |
| Memento | [Memento.go](BehavioralType/Memento.go) | [memento.md](BehavioralType/Docs/memento.md) | Save and restore state |
| Template Method | [Template.go](BehavioralType/Template.go) | [template-method.md](BehavioralType/Docs/template-method.md) | Fixed algorithm skeleton |
| State | [State.go](BehavioralType/State.go) | [state.md](BehavioralType/Docs/state.md) | State-driven behavior |
| Strategy | [Strategy.go](BehavioralType/Strategy.go) | [strategy.md](BehavioralType/Docs/strategy.md) | Swap algorithms |
| Observer | [Observer.go](BehavioralType/Observer.go) | [observer.md](BehavioralType/Docs/observer.md) | Event notification |
| Interpreter | [Interpreter.go](BehavioralType/Interpreter.go) | [interpreter.md](BehavioralType/Docs/interpreter.md) | Evaluate simple grammar |
| Chain of Responsibility | [Handler.go](BehavioralType/Handler.go) | [chain-of-responsibility.md](BehavioralType/Docs/chain-of-responsibility.md) | Pass requests along a chain |
| Iterator | [Iterator.go](BehavioralType/Iterator.go) | [iterator.md](BehavioralType/Docs/iterator.md) | Traverse collections uniformly |

## Design Principles

Read [solid-and-design-principles.md](solid-and-design-principles.md) before studying patterns. Patterns are useful only when they clarify responsibilities or reduce coupling.
