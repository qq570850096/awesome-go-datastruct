# Native HTTP Red Packet Server

## Goal

Run the red-packet service with only the Go standard library and a small unified handler wrapper.

## API

| Route | Method | Behavior |
| --- | --- | --- |
| `/redpacket/init` | POST | Initialize the pool with `total_amount` and `count`. |
| `/redpacket/grab` | POST | Grab one amount from the initialized pool. |

## Core Invariants

- Business handlers return data or an `appError`; the wrapper owns HTTP response formatting.
- Panic recovery, JSON writing, and logging stay centralized.
- Method validation should reject non-POST requests.

## Run Command

```bash
go run ./webdemo/redpacket/cmd/server
```

The server listens on `:8090`.

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8090/redpacket/init

curl -X POST http://localhost:8090/redpacket/grab
```

## Related Topics

- [redpacket](../../)
- [http_basic](../../../http_basic/)
- [minigin redpacket](../../../minigin/cmd/redpacket/)
