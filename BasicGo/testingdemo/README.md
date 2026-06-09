# Testing Demo

## Goal

Practice table-driven tests, error assertions, and benchmark shape with a small calculator function.

## Key Ideas

- Table-driven tests keep many cases readable.
- Error cases should be tested as first-class behavior.
- Benchmarks should isolate the operation being measured.
- Floating-point assertions should account for exactness or tolerance intentionally.

## Repository Code Map

| File | What to read for |
| --- | --- |
| calculator.go | `Calc`, supported operations, and calculation errors. |
| calculator_test.go | Table-driven tests and benchmark example. |

## Core Invariant

Each test case should name the behavior it protects and assert both result and error state.

## Practice Tasks

- Add a modulus-like operation for integers or document why it does not belong here.
- Add a divide-by-zero case if it is not already covered.
- Run the benchmark with `-benchmem` and explain the allocation result.

## Test Command

```bash
go test ./BasicGo/testingdemo
go test ./BasicGo/testingdemo -bench=. -benchmem
```

## Related Topics

- [basics](../basics/)
- [errors](../errors/)
