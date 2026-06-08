# Observer Pattern

Category: Behavioral

Code: [Observer.go](../Observer.go)

## Intent

Notify dependent objects automatically when a subject changes.

## Roles

- Subject manages subscriptions.
- Observer receives updates.
- Concrete subject publishes events.

## Use When

- Multiple listeners need to react to changes.
- The publisher should not know concrete subscribers.

## Tradeoffs

- Notification order can matter.
- Leaks are possible if subscribers are never removed.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestObserver
```
