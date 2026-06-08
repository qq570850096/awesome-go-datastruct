# MiniGin Demo

## Goal

Study a small HTTP router, route tree, context object, and middleware chain.

## Prerequisites

HTTP basics, trees, slices, and interfaces.

## Core Invariant

Route matching must return the correct handler and path parameters; middleware must run in registration order.

## Complexity

Route insertion and lookup depend on path depth and branching factor.

## Practice Tasks

- Add route conflict tests.
- Add middleware order tests.
- Add timeout and recovery behavior tests.

## Test Command

```bash
go test ./webdemo/minigin
```

## Related Topics

- [webdemo/http_basic](../http_basic/)
- [Graph_algo](../../Graph_algo/)
