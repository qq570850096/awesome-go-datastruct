# Errors

## Goal

Learn idiomatic error construction, wrapping, classification, and reporting with `errors.Is` and `errors.As`.

## Key Ideas

- Return errors for expected failure paths.
- Wrap errors with `%w` when callers need to inspect the original cause.
- Sentinel errors are useful for stable categories.
- Custom error types carry structured details for `errors.As`.

## Repository Code Map

| File | What to read for |
| --- | --- |
| errors.go | Sentinel error, validation error type, wrapping, and user-facing descriptions. |
| errors_test.go | Tests for wrapping and custom error extraction. |

## Core Invariant

Callers should be able to distinguish important error categories without parsing error strings.

## Practice Tasks

- Add another validation rule and assert it with `errors.As`.
- Add a new sentinel error and classify it with `errors.Is`.
- Rewrite one caller to return an error instead of a string description.

## Test Command

```bash
go test ./BasicGo/errors
```

## Related Topics

- [interface](../interface/)
- [testingdemo](../testingdemo/)
