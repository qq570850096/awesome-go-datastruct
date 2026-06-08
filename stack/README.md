# Stack

A stack restricts access to the most recent item. It models nested work such as call frames, undo stacks, and DFS.

## Learning Goals

- Use LIFO ordering intentionally.
- Connect stack behavior to recursion and DFS.
- Test empty and multi-item behavior.

## Prerequisites

- Slices and basic tests.

## Mental Model

A stack is a pile. Push adds to the top, pop removes from the top, and the middle is intentionally inaccessible.

## Diagram

```text
push 1, push 2, push 3
top -> 3 -> 2 -> 1
pop returns 3
```

## Terminology

| Term | Meaning |
| --- | --- |
| Top | The next item to be popped. |
| Push | Add an item to the top. |
| Pop | Remove and return the top item. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| stack.go | Stack operations. |
| stack_test.go | Order and boundary tests. |

## Core Invariants

- The last pushed item is popped first.
- Empty-pop behavior is explicit.

## Operation Walkthrough

Read push first: it appends to the backing storage. Read pop next: it checks emptiness, returns the last item, and shrinks storage.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Push | O(1) amortized | O(1) | Append to slice. |
| Pop | O(1) | O(1) | Remove last element. |

## Common Mistakes And Edge Cases

- Returning the bottom item by mistake.
- Not specifying empty-pop behavior.

## Worked Example

Push 1, 2, 3. Three pops return 3, 2, 1.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace push and pop on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty stack, one item, and repeated push/pop.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Push or Pop has the complexity shown in the table.

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
go test ./stack
```

## Next Topics

- [queue](../queue/)
- [Graph_algo/DFS](../Graph_algo/DFS/)
