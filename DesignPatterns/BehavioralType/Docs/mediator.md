# Mediator Pattern

Category: Behavioral

Code: [Mediator.go](../Mediator.go)

## Intent

Centralize communication between collaborating objects so they do not depend on each other directly.

## Roles

- Mediator defines communication operations.
- Concrete mediator coordinates object interactions.
- Colleagues send messages through the mediator.

## Use When

- Many objects need to coordinate through complex relationships.
- Direct dependencies are making change difficult.

## Tradeoffs

- The mediator can become too large.
- Keep coordination rules explicit and tested.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestMediator
```
