# Go Basics

## Goal

Learn the smallest building blocks used by every other module: variables, constants, control flow, functions, variadic arguments, named returns, type switches, and closures.

## Key Ideas

- Zero values make newly declared variables immediately usable.
- Constants are compile-time values and can participate in typed expressions.
- `if`, `for`, and `switch` are the core control-flow tools.
- Functions can accept variadic arguments, return named values, and close over local state.

## Repository Code Map

| File | What to read for |
| --- | --- |
| vars.go | Zero values, constants, and simple expressions. |
| control.go | `if`, `switch`, FizzBuzz, and type inspection. |
| funcs.go | Variadic sums, named returns, and closures. |
| basics_test.go | Behavior-focused tests for each small example. |

## Core Invariant

Each example should make one language rule observable through a small return value rather than relying on printed output.

## Practice Tasks

- Add a new table-driven case to `FizzBuzz`.
- Extend `TypeName` with another Go type.
- Write a second closure that tracks both count and sum.

## Test Command

```bash
go test ./BasicGo/basics
```

## Related Topics

- [pointers](../pointers/)
- [testingdemo](../testingdemo/)
