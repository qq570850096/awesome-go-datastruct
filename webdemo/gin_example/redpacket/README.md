# Gin Red Packet Demo

A Gin-powered HTTP entrypoint for the shared `webdemo/redpacket` business logic.

## Quick Start

```bash
go run ./webdemo/gin_example/redpacket
```

The server listens on `:8093`.

## What This Demo Shows

- How the same red-packet pool can be exposed through Gin.
- How `ShouldBindJSON` handles request decoding.
- How framework handlers stay thin when business logic lives in a shared package.
- How response shape stays consistent with the native HTTP and MiniGin demos.

## API

| Route | Method | Request body | Response data |
| --- | --- | --- | --- |
| `/redpacket/init` | POST | `{"total_amount":1000,"count":10}` | `{"ok":true}` |
| `/redpacket/grab` | POST | Empty body | `{"amount":15}` |

All responses use the same wrapper shape:

```json
{
  "code": 0,
  "msg": "ok",
  "data": {}
}
```

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8093/redpacket/init

curl -X POST http://localhost:8093/redpacket/grab
```

## Code Walkthrough

1. `gin.Default()` creates the engine with Gin's default logger and recovery middleware.
2. `redpacket.NewPool()` creates the shared in-memory pool.
3. `/redpacket/init` binds JSON into a small request struct and calls `pool.Init`.
4. `/redpacket/grab` calls `pool.Grab` and returns the amount.
5. `r.Run(":8093")` starts the demo server.

## Comparison With Other Entrypoints

| Entrypoint | Port | Main lesson |
| --- | --- | --- |
| [Native HTTP server](../../redpacket/cmd/server/) | `8090` | Manual wrapper, logging, recovery, and JSON encoding. |
| [MiniGin server](../../minigin/cmd/redpacket/) | `8092` | Local framework routing and middleware. |
| This Gin demo | `8093` | Framework binding and response helpers. |

## Common Pitfalls

- Putting allocation rules inside handlers instead of `webdemo/redpacket`.
- Forgetting to check `ShouldBindJSON` errors.
- Running multiple demos at once and sending requests to the wrong port.
- Treating this command package as a reusable library; it is an executable demo.

## Extension Ideas

- Add handler tests with `httptest` and a Gin router created in a helper.
- Add request validation for maximum packet count.
- Add a `GET /healthz` route for operational practice.
- Compare Gin middleware behavior with MiniGin middleware behavior.

## Related Topics

- [gin_example](../)
- [redpacket](../../redpacket/)
- [minigin redpacket](../../minigin/cmd/redpacket/)
