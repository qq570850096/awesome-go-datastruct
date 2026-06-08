# Facade Pattern

Category: Structural

Code: [Facade.go](../Facade.go)

## Intent

Provide a simple entrypoint over a more complex subsystem.

## Roles

- Facade exposes a small API.
- Subsystem types do the detailed work.
- Client calls the facade instead of coordinating many objects.

## Use When

- A workflow requires several subsystem calls.
- Callers need a stable and simpler API.

## Tradeoffs

- Facade can become a large coordinator.
- It should not hide every useful subsystem capability.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestFacade
```
