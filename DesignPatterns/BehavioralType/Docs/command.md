# Command Pattern

Category: Behavioral

Code: [Command.go](../Command.go)

## Intent

Wrap a request as an object so it can be queued, logged, retried, or invoked later.

## Roles

- Command defines the execution method.
- Concrete command binds a receiver and action.
- Invoker triggers commands.
- Receiver performs the actual work.

## Use When

- Requests need to be stored or replayed.
- The invoker should not know receiver details.

## Tradeoffs

- Many operations can create many command types.
- Keep command objects small.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestCommand
```
