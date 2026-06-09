# Gin Red Packet Demo

## Goal

Expose the shared red-packet pool through Gin so the same business logic can be compared against native HTTP and MiniGin implementations.

## API

| Route | Method | Behavior |
| --- | --- | --- |
| `/redpacket/init` | POST | Initialize the pool with `total_amount` and `count`. |
| `/redpacket/grab` | POST | Grab one amount from the initialized pool. |

## Core Invariants

- Handler code should stay thin and delegate red-packet rules to `webdemo/redpacket`.
- JSON responses use `{code, msg, data}` consistently.
- Invalid input should return HTTP 400 with a useful message.

## Run Command

```bash
go run ./webdemo/gin_example/redpacket
```

The demo listens on `:8093`.

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8093/redpacket/init

curl -X POST http://localhost:8093/redpacket/grab
```

## Related Topics

- [gin_example](../)
- [redpacket](../../redpacket/)
- [minigin redpacket](../../minigin/cmd/redpacket/)
