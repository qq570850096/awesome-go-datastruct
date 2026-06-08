# Singleton Pattern

Category: Creational

Code: [Singleton.go](../Singleton.go)

## Intent

Guarantee that a type has one shared instance and provide a controlled access point.

## Roles

- Singleton stores the shared instance.
- Constructor or accessor returns the same instance.
- Synchronization protects lazy initialization when needed.

## Use When

- A resource must be shared globally.
- Instance creation is expensive and should happen once.

## Tradeoffs

- Global state can make tests harder.
- Prefer dependency injection when callers need flexibility.

## Test Command

```bash
go test ./DesignPatterns/CreativeType -run TestSingleton
```
