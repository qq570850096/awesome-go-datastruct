# Errors

Idiomatic error construction, wrapping, classification, and reporting with sentinel errors and custom error types.

## Quick Start

```bash
go test ./BasicGo/errors
go test ./BasicGo/errors -run TestLoadConfigErrorWrapping
```

## What You Will Learn

- How sentinel errors describe stable failure categories.
- How `%w` preserves an error chain for `errors.Is`.
- How custom error types carry structured details for `errors.As`.
- How user-facing descriptions can be separated from machine-readable errors.

## Concept Map

```text
operation fails
      |
      +-- sentinel error -> errors.Is
      |
      +-- custom type    -> errors.As
      |
      +-- wrapped error  -> keeps original cause
```

## API Surface

| Function or type | Purpose | Important contract |
| --- | --- | --- |
| `ErrConfigNotFound` | Stable config-missing category. | Callers inspect it with `errors.Is`. |
| `ValidationError` | Carries `Field` and `Reason`. | Callers inspect it with `errors.As`. |
| `LoadConfig(path)` | Simulate config loading. | Wraps `ErrConfigNotFound` when path is empty. |
| `ValidateUser(name, age)` | Validate user input. | Returns `ValidationError` for invalid fields. |
| `HandleConfig(path)` | Convert config errors into descriptions. | Demonstrates classification. |
| `ValidateAndDescribe(name, age)` | Convert validation errors into descriptions. | Demonstrates custom error extraction. |

## Guided Walkthrough

1. Read the declaration of `ErrConfigNotFound`.
2. Follow `LoadConfig` and identify where `%w` preserves the sentinel.
3. Read `ValidationError` and its `Error()` method.
4. Finish with the tests; they are the contract for `errors.Is` and `errors.As`.

## Example

```go
err := LoadConfig("")
if errors.Is(err, ErrConfigNotFound) {
    fmt.Println("use default config")
}
```

The displayed message can change, but the sentinel category remains stable.

## Common Pitfalls

- Parsing error strings instead of using `errors.Is` or `errors.As`.
- Wrapping with `%v` when callers need `%w`.
- Returning custom error pointers in one place and values in another without testing `errors.As`.
- Turning every validation failure into a separate sentinel error.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add a validation case for an invalid email-like field. |
| Drill | Add a second sentinel error and classify it in `HandleConfig`. |
| Challenge | Replace `ValidateAndDescribe` with an error-returning API and move formatting to the caller. |

## Quality Checklist

- Important categories are inspectable without string matching.
- Tests cover both success and failure paths.
- Wrapped errors preserve the cause intentionally.
- User-facing text is not the only source of truth.

## Related Topics

- [interface](../interface/)
- [testingdemo](../testingdemo/)
- [webdemo/http_basic](../../webdemo/http_basic/)
