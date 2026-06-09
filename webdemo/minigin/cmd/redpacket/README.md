# MiniGin Red Packet Demo

An executable red-packet service built on the local MiniGin framework.

## Quick Start

```bash
go run ./webdemo/minigin/cmd/redpacket
```

The server listens on `:8092`.

## What This Demo Shows

- How a small custom framework handles route registration.
- How MiniGin middleware composes recovery and logging.
- How `Context.BindJSON` and `Context.JSON` create a framework-style handler experience.
- How reusable business logic stays in `webdemo/redpacket`.

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

## Example Requests

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":10}' \
  http://localhost:8092/redpacket/init

curl -X POST http://localhost:8092/redpacket/grab
```

## Code Walkthrough

1. `minigin.New()` creates the local framework engine.
2. `engine.Use(minigin.Recovery(), minigin.Logger())` attaches middleware.
3. `engine.POST("/redpacket/init", ...)` registers JSON initialization.
4. `engine.POST("/redpacket/grab", ...)` registers pool grabbing.
5. `http.ListenAndServe(":8092", engine)` proves MiniGin implements `http.Handler`.

## Comparison With Other Entrypoints

| Entrypoint | Port | Main lesson |
| --- | --- | --- |
| [Native HTTP server](../../../redpacket/cmd/server/) | `8090` | Standard-library wrapper and middleware. |
| This MiniGin demo | `8092` | Custom router, context, and middleware chain. |
| [Gin demo](../../../gin_example/redpacket/) | `8093` | Production framework ergonomics. |

## Common Pitfalls

- Forgetting middleware order when debugging response behavior.
- Writing direct `net/http` responses inside MiniGin handlers instead of using `Context.JSON`.
- Reimplementing red-packet allocation inside this command package.
- Ignoring the difference between framework context and `context.Context`.

## Extension Ideas

- Add route tests under `webdemo/minigin` for the same path patterns.
- Add authentication middleware and compare it with the native wrapper.
- Add request IDs to the MiniGin context.
- Add a graceful shutdown example.

## Related Topics

- [minigin](../../)
- [redpacket](../../../redpacket/)
- [Gin redpacket](../../../gin_example/redpacket/)
