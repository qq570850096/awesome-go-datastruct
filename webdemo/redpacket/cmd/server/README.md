# Native HTTP Red Packet Server

A standard-library HTTP server for the shared red-packet pool, designed to show explicit routing, wrapping, JSON responses, logging, and panic recovery.

## Quick Start

```bash
go run ./webdemo/redpacket/cmd/server
```

The server listens on `:8090`.

## What This Demo Shows

- How to build a small service with `net/http` only.
- How a wrapper separates business handlers from HTTP response formatting.
- How to centralize panic recovery and JSON encoding.
- How to keep request parsing and domain rules in clear layers.

## API

| Route | Method | Request body | Response data |
| --- | --- | --- | --- |
| `/redpacket/init` | POST | `{"total_amount":1000,"count":10}` | `{"ok":true}` |
| `/redpacket/grab` | POST | Empty body | `{"amount":15}` |

All responses use this shape:

```json
{
  "code": 0,
  "msg": "ok",
  "data": {}
}
```

Errors use the same shape with a non-zero `code` and a message.

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8090/redpacket/init

curl -X POST http://localhost:8090/redpacket/grab
```

## Code Walkthrough

1. `redpacket.NewPool()` creates the shared pool.
2. `http.NewServeMux()` registers two wrapped handlers.
3. `poolInitHandler` parses JSON and calls `pool.Init`.
4. `poolGrabHandler` calls `pool.Grab`.
5. `wrap` handles panic recovery, application errors, and JSON output.
6. `loggingMiddleware` records method, path, and elapsed time.
7. `http.Server` sets read and write timeouts before listening.

## Handler Shape

```go
type appHandler func(r *http.Request) (interface{}, *appError)
```

Business handlers return data or an application error. The wrapper owns HTTP status codes, headers, response encoding, and recovery. This keeps route logic small and consistent.

## Comparison With Other Entrypoints

| Entrypoint | Port | Main lesson |
| --- | --- | --- |
| This native server | `8090` | Explicit standard-library HTTP structure. |
| [MiniGin server](../../../minigin/cmd/redpacket/) | `8092` | Custom framework abstraction over `net/http`. |
| [Gin demo](../../../gin_example/redpacket/) | `8093` | Third-party framework ergonomics. |

## Common Pitfalls

- Writing multiple responses for one request.
- Forgetting to set `Content-Type` before writing JSON.
- Returning HTTP 200 for application errors.
- Putting panic recovery in every handler instead of centralizing it.
- Sharing mutable state without understanding the red-packet pool's locking behavior.

## Extension Ideas

- Add `httptest` coverage for invalid method, invalid JSON, and pool exhaustion.
- Add graceful shutdown using `context.Context`.
- Add request IDs and include them in logs.
- Extract router construction into a testable helper.

## Related Topics

- [redpacket](../../)
- [http_basic](../../../http_basic/)
- [minigin redpacket](../../../minigin/cmd/redpacket/)
