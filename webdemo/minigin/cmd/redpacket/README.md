# MiniGin Red Packet Demo

## Goal

Run the red-packet service on the local MiniGin framework to exercise routing, middleware, JSON binding, and shared business logic.

## API

| Route | Method | Behavior |
| --- | --- | --- |
| `/redpacket/init` | POST | Initialize the pool with `total_amount` and `count`. |
| `/redpacket/grab` | POST | Grab one amount from the initialized pool. |

## Core Invariants

- MiniGin middleware should run before handlers and recover from panics.
- Handlers should use `BindJSON` and `JSON` rather than direct `net/http` response writes.
- Red-packet allocation rules stay in `webdemo/redpacket`.

## Run Command

```bash
go run ./webdemo/minigin/cmd/redpacket
```

The demo listens on `:8092`.

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8092/redpacket/init

curl -X POST http://localhost:8092/redpacket/grab
```

## Related Topics

- [minigin](../../)
- [redpacket](../../../redpacket/)
- [Gin redpacket](../../../gin_example/redpacket/)
