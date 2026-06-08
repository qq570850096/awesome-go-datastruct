# Circular Queue

A circular queue teaches fixed-capacity design and modulo arithmetic. It is a useful bridge from a simple queue to ring buffers.

## Learning Goals

- Model wraparound indexes.
- Distinguish full and empty states.
- Test capacity boundaries.

## Prerequisites

- Arrays, indexes, and FIFO queues.

## Mental Model

A circular queue reuses fixed storage. Head and tail move forward modulo capacity instead of shifting elements.

## Diagram

```text
indexes: 0 1 2 3
values:  A B _ _
head=0 tail=2
next = (index + 1) % capacity
```

## Terminology

| Term | Meaning |
| --- | --- |
| Capacity | Maximum number of elements. |
| Head | Index of the front element. |
| Tail | Index where the next element may be inserted. |
| Wraparound | Moving from the last slot back to index 0. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| 622.go | Circular queue implementation. |
| README.md | This textbook chapter. |

## Core Invariants

- Head and tail stay within `[0, capacity)`.
- Full and empty states are never confused.

## Operation Walkthrough

Trace enqueue until the tail reaches the end, then apply modulo arithmetic. Trace dequeue separately and verify the front item changes without moving stored values.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| EnQueue | O(1) | O(1) | One write and index update. |
| DeQueue | O(1) | O(1) | One index update. |
| Front/Rear | O(1) | O(1) | Direct indexed access. |

## Common Mistakes And Edge Cases

- Using modulo on one index but not the other.
- Forgetting capacity-one behavior.
- Using the same state for full and empty.

## Worked Example

With capacity 3, enqueue A, B, C, dequeue A, then enqueue D. D lands in the slot previously used by A.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace wraparound after dequeue on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for capacity one, full queue, empty queue, and wraparound.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why EnQueue or DeQueue has the complexity shown in the table.

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
go test ./main
```

## Next Topics

- [queue](../queue/)
- [Graph_algo/BFS](../Graph_algo/BFS/)
