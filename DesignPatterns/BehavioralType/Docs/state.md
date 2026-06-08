# State Pattern

Category: Behavioral

Code: [State.go](../State.go)

## Intent

Let an object change behavior when its internal state changes.

## Roles

- Context stores current state.
- State interface defines state-dependent behavior.
- Concrete states implement behavior and transitions.

## Use When

- Large conditional blocks depend on object state.
- New states should be added without rewriting the whole context.

## Tradeoffs

- Many states create many types.
- Transitions must be easy to trace in tests.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestState
```
