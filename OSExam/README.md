# Operating System Exercises

## Goal

Practice scheduling and file-system modeling with small Go examples.

## Prerequisites

Sorting, state modeling, and tests.

## Core Invariant

Scheduling examples must compute wait time, turnaround time, and weighted turnaround time consistently.

## Complexity

Varies by scheduling algorithm; most examples scan process lists and are O(n^2) or better depending on selection strategy.

## Practice Tasks

- Add edge cases for equal arrival times.
- Add tests for empty and single-process input.
- Separate demo output from assertive tests where practical.

## Test Command

```bash
go test ./OSExam
```

## Related Topics

- [Sorts](../Sorts/)
- [queue](../queue/)
