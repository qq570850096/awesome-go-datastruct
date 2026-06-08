# Queue

A queue restricts access to the oldest item. It is the core model for waiting lines, BFS, and scheduling.

## Learning Goals

- Use FIFO ordering intentionally.
- Understand head and tail movement.
- Connect queue behavior to BFS.

## Prerequisites

- Slices and basic tests.

## Mental Model

A queue is a line. Enqueue at the back, dequeue from the front, and preserve arrival order.

## Diagram

```text
front -> 1 -> 2 -> 3 <- back
dequeue returns 1
```

## Terminology

| Term | Meaning |
| --- | --- |
| Front | The next item to leave. |
| Back | Where new items enter. |
| FIFO | First in, first out. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| queue.go | Queue operations. |
| queue_test.go | FIFO behavior tests. |

## Core Invariants

- The earliest enqueued item is dequeued first.
- Empty-dequeue behavior is explicit.

## Operation Walkthrough

Read enqueue as adding work to the back. Read dequeue as removing from the front. Then compare the implementation with BFS, where each discovered vertex waits for its turn.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Enqueue | O(1) amortized | O(1) | Append to storage. |
| Dequeue | O(1) or O(n) | O(1) | Depends on whether the implementation shifts or advances a head index. |

## Common Mistakes And Edge Cases

- Accidentally building stack behavior.
- Retaining old front elements in long-lived queues.

## Worked Example

Enqueue 1, 2, 3. Three dequeues return 1, 2, 3.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace enqueue and dequeue on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty queue, one item, and alternating enqueue/dequeue.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Enqueue or Dequeue has the complexity shown in the table.

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
go test ./queue
```

## Next Topics

- [stack](../stack/)
- [Graph_algo/BFS](../Graph_algo/BFS/)
