# Behavioral Design Patterns

Behavioral patterns describe how objects communicate, delegate work, vary algorithms, react to state changes, and traverse collections.

## Quick Start

```bash
go test ./DesignPatterns/BehavioralType
```

Run one pattern while studying it:

```bash
go test ./DesignPatterns/BehavioralType -run TestConcreteCommand
go test ./DesignPatterns/BehavioralType -run TestTelev
```

## What This Module Covers

- Request encapsulation with command objects.
- Handler chains for escalating requests.
- State-driven behavior changes.
- Runtime algorithm selection.
- Subscriber notification and mediator coordination.
- Snapshot restoration, iteration, interpretation, and template-method hooks.

## Pattern Index

| Pattern | Code | Detailed doc | Primary lesson |
| --- | --- | --- | --- |
| Command | [Command.go](Command.go) | [command.md](Docs/command.md) | An invoker can trigger an action without knowing receiver details. |
| Chain of Responsibility | [Handler.go](Handler.go) | [chain-of-responsibility.md](Docs/chain-of-responsibility.md) | A request moves through handlers until one is responsible. |
| Interpreter | [Interpreter.go](Interpreter.go) | [interpreter.md](Docs/interpreter.md) | A grammar can be represented as expression objects. |
| Iterator | [Iterator.go](Iterator.go) | [iterator.md](Docs/iterator.md) | A collection exposes traversal without leaking storage details. |
| Mediator | [Mediator.go](Mediator.go) | [mediator.md](Docs/mediator.md) | Colleagues communicate through a coordinating object. |
| Memento | [Memento.go](Memento.go) | [memento.md](Docs/memento.md) | State can be captured and restored without exposing internals. |
| Observer | [Observer.go](Observer.go) | [observer.md](Docs/observer.md) | Subscribers are notified when a subject changes. |
| State | [State.go](State.go) | [state.md](Docs/state.md) | Behavior changes when the current state object changes. |
| Strategy | [Strategy.go](Strategy.go) | [strategy.md](Docs/strategy.md) | Algorithms are swapped behind stable behavior interfaces. |
| Template Method | [Template.go](Template.go) | [template-method.md](Docs/template-method.md) | A fixed skeleton calls overridable steps. |

## Learning Path

1. Start with [Strategy.go](Strategy.go). It is the easiest way to see "behavior as a pluggable object".
2. Move to [State.go](State.go). Compare it with strategy: state changes itself over time, while strategy is usually chosen by the client.
3. Read [Command.go](Command.go) and [Handler.go](Handler.go) to compare request-as-object with request-as-chain.
4. Study [Observer.go](Observer.go) and [Mediator.go](Mediator.go) together; both reduce direct coupling, but with different control flow.
5. Finish with [Interpreter.go](Interpreter.go), [Iterator.go](Iterator.go), [Memento.go](Memento.go), and [Template.go](Template.go).

## Role Vocabulary

| Role | Meaning in this module |
| --- | --- |
| Client | Code that wires concrete collaborators together. |
| Invoker | Object that triggers a command or workflow. |
| Receiver | Object that performs real work for a command. |
| Handler | Object that may handle or forward a request. |
| Subject | Object that owns changing state and notifies observers. |
| Strategy/state object | Object that holds one behavior variant. |

## Design Notes

- Keep pattern examples small. The point is the responsibility boundary, not domain complexity.
- Prefer interfaces at collaboration boundaries and concrete types for simple value holders.
- Tests should prove the collaboration path: command reaches receiver, observer gets notified, state transitions occur, iterator visits expected elements.
- Avoid using a pattern name as a reason to add indirection. The pattern should solve a visible coupling or variation problem.

## Common Pitfalls

- Confusing strategy and state because both delegate behavior.
- Letting a mediator become a god object that owns too much business logic.
- Creating one command type per trivial function when a function value would be enough.
- Testing only that methods can be called, not that collaboration changed the right state.
- Hiding invalid state transitions instead of asserting them.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add one boundary test for an existing pattern, such as an empty iterator or missing observer. |
| Drill | Draw the collaboration diagram for command, observer, and state. |
| Challenge | Implement an undoable command using the memento module as support. |
| Extension | Add a new behavioral pattern doc under `Docs/` and link it from this index. |

## Quality Checklist

- Every pattern has code, tests, and a linked detailed doc.
- Each example identifies its roles clearly.
- Tests assert behavior, not only printed output.
- New patterns are linked from [DesignPatterns README](../README.MD) when added.

## Related Topics

- [DesignPatterns](../)
- [CreativeType](../CreativeType/)
- [StructuralType](../StructuralType/)
- [BasicGo/interface](../../BasicGo/interface/)
