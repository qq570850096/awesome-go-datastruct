# Go Fundamentals

This chapter is the foundation for the rest of the repository. It explains the Go mechanics that later data structures rely on: values, pointers, slices, maps, methods, interfaces, errors, tests, and concurrency.

## Learning Goals

- Read small Go packages as executable lessons.
- Predict when data is copied and when state is shared.
- Use tests as the feedback loop for language learning.
- Recognize concurrency primitives and their engineering purpose.

## Prerequisites

- A working Go toolchain.
- Comfort running `go test` from the repository root.

## Mental Model

Go programs are small packages connected by explicit imports. Ordinary values are copied, while pointers, slices, maps, channels, and interfaces carry references or runtime descriptors. When in doubt, write a test and observe the caller-visible behavior.

## Diagram

```text
value assignment -> copy
&value           -> address
slice            -> descriptor + backing array
map              -> reference to hash table
interface        -> dynamic type + dynamic value
context          -> cancellation + deadline + request values
```

## Terminology

| Term | Meaning |
| --- | --- |
| Zero value | The default usable value of a type. |
| Pointer | A value that stores an address. |
| Receiver | The binding parameter for a method. |
| Interface | An implicitly satisfied behavior contract. |
| Context | A request-scoped cancellation and value carrier. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| basics/ | Variables, constants, control flow, functions, and closures. |
| pointers/ | Address semantics and caller-visible mutation. |
| structs/, interface/ | Data modeling and behavior abstraction. |
| slicemap/ | Slice backing arrays, maps, and set patterns. |
| GoRoutine/, channelselect/, context/, sharedvars/ | Concurrency, communication, cancellation, and synchronization. |
| testingdemo/ | Table-driven tests and benchmark shape. |

## Core Invariants

- Tests should assert behavior rather than only print output.
- Pointer receivers are required for caller-visible mutation.
- Slices may share backing arrays until growth allocates a new array.
- Shared mutable state needs synchronization or a design that avoids sharing.

## Operation Walkthrough

Study `basics` first, then `pointers`, then `structs` and `interface`. After language basics, read concurrency in this order: start work with goroutines, communicate with channels, cancel with context, and protect shared variables with synchronization.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Map lookup | Average O(1) | O(1) | Hash-table lookup is expected constant time. |
| Slice append | Amortized O(1) | O(n) when growing | Growth may copy existing elements. |
| Table test run | O(cases) | O(1) | Each case runs once plus tested code cost. |

## Common Mistakes And Edge Cases

- Assuming every assignment shares state.
- Ignoring slice backing-array sharing.
- Using `panic` for ordinary input failures.
- Using sleeps where deterministic synchronization would be better.

## Worked Example

A function that accepts `int` receives a copy, so a swap cannot modify the caller variables. A function that accepts `*int` can modify the caller-visible memory. The same idea powers linked lists and trees.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace pointer mutation and a table-driven test on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for zero values, nil pointers, empty slices, and canceled contexts.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why a map lookup or pointer update has the complexity shown in the table.

<details>
<summary>Hint</summary>

Count the number of nodes, array cells, characters, or edges that can be visited. Then count extra storage.

</details>

<details>
<summary>Reference answer</summary>

A good answer separates input size from auxiliary state. It mentions whether the operation follows one path, scans all elements, visits all edges, or allocates a helper structure.

</details>

## Test And Benchmark Commands

```bash
go test ./BasicGo/...
go test -race ./BasicGo/sharedvars
go test ./BasicGo/testingdemo -bench=.
```

## Next Topics

- [Linked](../Linked/) for pointer-heavy practice.
- [Graph algorithms](../Graph_algo/) for traversal state.
