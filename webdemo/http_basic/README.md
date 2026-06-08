# Native HTTP Demo

## Goal

Learn a minimal REST-style service with the Go standard library.

## Prerequisites

Basic Go, errors, JSON, and interfaces.

## Core Invariant

Handlers should parse requests, validate input, write one response, and keep shared state synchronized.

## Complexity

Request handling is usually O(1) plus any data-store lookup cost in this demo.

## Practice Tasks

- Add tests for bad JSON.
- Add tests for missing routes.
- Trace middleware execution order.

## Test Command

```bash
go test ./webdemo/http_basic
```

## Related Topics

- [webdemo/minigin](../minigin/)
- [BasicGo/context](../../BasicGo/context/)
