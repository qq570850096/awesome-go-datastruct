# Behavioral Design Patterns

## Goal

Study patterns that organize object communication, algorithm variation, state changes, request handling, and traversal.

## Pattern Map

| Pattern | Code | Detailed doc | Core idea |
| --- | --- | --- | --- |
| Command | [Command.go](Command.go) | [command.md](Docs/command.md) | Wrap a request as an executable object. |
| Chain of Responsibility | [Handler.go](Handler.go) | [chain-of-responsibility.md](Docs/chain-of-responsibility.md) | Pass a request through handlers until one handles it. |
| Interpreter | [Interpreter.go](Interpreter.go) | [interpreter.md](Docs/interpreter.md) | Evaluate a small grammar with expression objects. |
| Iterator | [Iterator.go](Iterator.go) | [iterator.md](Docs/iterator.md) | Traverse a collection through a stable interface. |
| Mediator | [Mediator.go](Mediator.go) | [mediator.md](Docs/mediator.md) | Coordinate colleagues through a central object. |
| Memento | [Memento.go](Memento.go) | [memento.md](Docs/memento.md) | Capture and restore state snapshots. |
| Observer | [Observer.go](Observer.go) | [observer.md](Docs/observer.md) | Notify subscribers when state changes. |
| State | [State.go](State.go) | [state.md](Docs/state.md) | Let behavior vary with internal state. |
| Strategy | [Strategy.go](Strategy.go) | [strategy.md](Docs/strategy.md) | Swap algorithms behind a common interface. |
| Template Method | [Template.go](Template.go) | [template-method.md](Docs/template-method.md) | Keep an algorithm skeleton fixed while steps vary. |

## Core Invariants

- The caller should depend on a role interface, not on concrete collaborators.
- State transitions and handler chains should be visible in tests.
- Pattern examples should stay small enough to reveal the responsibility boundary.

## Practice Tasks

- Add a negative or boundary test for one pattern.
- Draw the collaboration diagram for command, observer, or state.
- Compare strategy and state: both vary behavior, but for different reasons.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType
```

## Related Topics

- [DesignPatterns](../)
- [CreativeType](../CreativeType/)
- [StructuralType](../StructuralType/)
