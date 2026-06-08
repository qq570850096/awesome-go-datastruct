# Proxy Pattern

Category: Structural

Code: [Proxy.go](../Proxy.go)

## Intent

Control access to another object through a stand-in that implements the same interface.

## Roles

- Subject defines the interface.
- Real subject performs the work.
- Proxy controls access, caching, logging, or lazy creation.

## Use When

- Access needs validation, caching, lazy loading, or remote handling.
- Callers should use the same interface as the real object.

## Tradeoffs

- Proxy behavior can surprise callers if it is not documented.
- Keep policy separate from core work when possible.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestProxy
```
