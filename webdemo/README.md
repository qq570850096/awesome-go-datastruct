# Web Demos

This directory contains small Web examples that connect Go language mechanics to service-style code.

| Path | Focus | Validation |
| --- | --- | --- |
| [http_basic](http_basic/) | Native `net/http`, handlers, middleware, and REST-style routing | `go test ./webdemo/http_basic` |
| [minigin](minigin/) | A tiny router and middleware framework inspired by Gin | `go test ./webdemo/minigin` |
| [gin_example](gin_example/) | A small Gin usage example | Manual demo |

## Learning Goals

- Understand request and response boundaries.
- Practice middleware composition.
- Connect `context`, errors, and tests to HTTP code.
- Keep demos small enough for learning.
