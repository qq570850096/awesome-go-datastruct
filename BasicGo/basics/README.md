# Go Basics

Small, testable examples for the syntax and control-flow rules that every later module relies on.

## Quick Start

```bash
go test ./BasicGo/basics
```

Run a single behavior group:

```bash
go test ./BasicGo/basics -run TestControlFlow
```

## What You Will Learn

- How zero values make variables usable immediately.
- How constants, expressions, and ordinary functions compose.
- How `if`, `for`, and `switch` encode branching rules.
- How variadic arguments, named returns, and closures appear in real code.
- How tiny examples become stronger when they are protected by table-driven tests.

## Concept Map

```text
declaration -> zero value
constant    -> compile-time value
if/for      -> branching and repetition
switch      -> value or type dispatch
function    -> inputs + return values
closure     -> function + captured state
test        -> executable specification
```

## API Surface

| Function or value | File | Purpose |
| --- | --- | --- |
| `ZeroValues()` | vars.go | Return default values for `int`, `string`, and `bool`. |
| `Pi`, `DoublePi()` | vars.go | Show a constant and a simple derived expression. |
| `Max(a, b)` | control.go | Minimal `if` branch. |
| `FizzBuzz(n)` | control.go | Branch ordering and divisibility checks. |
| `TypeName(v)` | control.go | Type switch over `any`. |
| `Sum(nums...)` | funcs.go | Variadic parameters. |
| `SplitName(full)` | funcs.go | Named return values. |
| `NewCounter(start)` | funcs.go | Closure that captures mutable state. |

## Reading Guide

1. Start with `vars.go` and make sure you can predict every returned zero value.
2. Move to `control.go`; read `FizzBuzz` before `TypeName` so value switches and type switches stay separate.
3. Finish with `funcs.go`; closures are the bridge from basic syntax to stateful structures.
4. Read `basics_test.go` last and notice how the tests describe behavior without printing.

## Example

```go
counter := NewCounter(10)
fmt.Println(counter()) // 11
fmt.Println(counter()) // 12
```

The returned function keeps access to `start` after `NewCounter` returns. That idea reappears later in iterators, callbacks, and middleware.

## Common Pitfalls

- Checking `FizzBuzz` conditions in the wrong order and returning `Fizz` for `15`.
- Assuming a named return value means a function should avoid explicit `return value` statements.
- Forgetting that closures share captured variables, not a frozen copy of every value.
- Using `any` before a concrete type or small interface would be clearer.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add table cases for `FizzBuzz(0)`, `FizzBuzz(30)`, and a prime number. |
| Drill | Extend `TypeName` with `[]int` and explain why it belongs in the type switch. |
| Challenge | Write `NewAccumulator(start)` that returns both `add(delta)` and `value()` closures. |

## Quality Checklist

- Tests cover normal, boundary, and surprising inputs.
- Examples return values instead of relying only on `fmt.Println`.
- New functions stay small enough that the language rule remains visible.

## Related Topics

- [pointers](../pointers/)
- [testingdemo](../testingdemo/)
- [slicemap](../slicemap/)
