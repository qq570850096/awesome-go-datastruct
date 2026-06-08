# Chain Of Responsibility Pattern

Category: Behavioral

Code: [Handler.go](../Handler.go)

## Intent

Pass a request along a chain until one handler processes it.

## Roles

- Handler defines request handling and next link.
- Concrete handlers decide whether they can handle the request.
- Client builds the chain.

## Use When

- Several handlers may process a request.
- The sender should not choose a concrete handler directly.

## Tradeoffs

- A request may go unhandled.
- Long chains can be hard to debug.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestHandler
```
