# Memento Pattern

Category: Behavioral

Code: [Memento.go](../Memento.go)

## Intent

Capture and restore an object state without exposing its internal representation.

## Roles

- Originator creates and restores snapshots.
- Memento stores snapshot state.
- Caretaker manages snapshots without inspecting them.

## Use When

- Undo or rollback behavior is needed.
- State should be restored without breaking encapsulation.

## Tradeoffs

- Snapshots can consume memory.
- Snapshot boundaries must be explicit.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestMemento
```
