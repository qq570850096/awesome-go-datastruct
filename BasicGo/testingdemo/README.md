# Testing Demo

Table-driven tests, named subtests, error assertions, and benchmark shape around a small calculator.

## Quick Start

```bash
go test ./BasicGo/testingdemo
go test ./BasicGo/testingdemo -bench=. -benchmem
```

Run one subtest group:

```bash
go test ./BasicGo/testingdemo -run 'TestCalc/div-by-zero'
```

## What You Will Learn

- How table-driven tests keep many cases compact.
- How named subtests make failures easy to locate.
- How to assert both results and error categories.
- How benchmark loops use `b.N`.
- Why benchmark setup should not dominate the measured operation.

## Concept Map

```text
function contract -> test table
case name         -> subtest
expected result   -> assertion
expected error    -> assertion
benchmark         -> repeat operation b.N times
```

## API Surface

| Function or value | Purpose | Important contract |
| --- | --- | --- |
| `Operation` | String-backed operation selector. | Supported values are `add`, `sub`, `mul`, and `div`. |
| `Add`, `Sub`, `Mul`, `Div` | Known operations. | Used by `Calc`. |
| `ErrUnknownOp` | Unsupported operation category. | Returned for unknown operations. |
| `ErrDivideByZero` | Division-by-zero category. | Returned when `Div` receives zero divisor. |
| `Calc(a, b, op)` | Perform one calculator operation. | Returns result plus error. |
| `TestCalc` | Table-driven behavior coverage. | Uses subtests with case names. |
| `BenchmarkCalc` | Benchmark example. | Measures repeated calculator calls. |

## Guided Walkthrough

1. Read `Calc` and write down every branch.
2. Read the test table and verify that every branch has at least one case.
3. Notice how errors are compared directly because they are sentinel values.
4. Read `BenchmarkCalc`; identify setup work versus measured work.

## Example

```go
got, err := Calc(10, 0, Div)
if err == ErrDivideByZero {
    fmt.Println("invalid divisor")
}
fmt.Println(got) // 0
```

The caller gets a normal error instead of a panic.

## Common Pitfalls

- Testing only successful cases.
- Forgetting to name table cases clearly.
- Comparing floating-point values without thinking about precision.
- Doing expensive random setup inside a benchmark loop when it is not part of the target behavior.

## Exercises

| Level | Task |
| --- | --- |
| Warm-up | Add cases for negative operands. |
| Drill | Replace direct error comparison with `errors.Is` and explain whether it changes behavior. |
| Challenge | Add a benchmark that measures one fixed operation without random selection. |

## Quality Checklist

- Every branch has a test case.
- Error behavior is asserted explicitly.
- Subtest names describe behavior.
- Benchmark setup is intentional.

## Related Topics

- [basics](../basics/)
- [errors](../errors/)
- [Sorts](../../Sorts/)
