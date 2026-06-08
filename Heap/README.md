# Heap And Priority Queue

A heap stores partial order in an array. It gives fast access to the highest-priority item without fully sorting all items.

## Learning Goals

- Map tree positions to array indexes.
- Maintain heap order with sift-up and sift-down.
- Use heaps for priority queues.

## Prerequisites

- Arrays, comparisons, and tree vocabulary.

## Mental Model

A heap is a complete tree laid out in an array. Parent-child relationships come from indexes, not pointers.

## Diagram

```text
index: 0 1 2 3 4
value: 9 7 5 1 3
parent(i) = (i-1)/2
left(i) = 2*i+1
right(i) = 2*i+2
```

## Terminology

| Term | Meaning |
| --- | --- |
| Heap order | Parent priority is at least child priority in a max-heap. |
| Sift up | Move a new value toward the root. |
| Sift down | Move a replacement root down. |
| Heapify | Build a heap from an array. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Array.go | Array-backed heap operations. |
| PriorityQueue.go | Priority queue shape. |
| Arrar_test.go | Heap behavior checks. |

## Core Invariants

- The root is the maximum in a max-heap.
- Every level is filled left to right except possibly the last.

## Operation Walkthrough

Insertion appends the new value, then sifts it up while it outranks its parent. Removal swaps or replaces the root, then sifts down by choosing the larger child until heap order returns.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Peek max | O(1) | O(1) | Root is index 0. |
| Insert | O(log n) | O(1) | Moves up tree height. |
| Remove max | O(log n) | O(1) | Moves down tree height. |
| Heapify | O(n) | O(1) | Bottom-up repair is linear. |

## Common Mistakes And Edge Cases

- Expecting the whole array to be sorted.
- Forgetting to compare both children in sift-down.
- Calling parent on index 0.

## Worked Example

Insert 9, 7, 5, 10. The 10 is appended, compared with parents, and sifted to the root.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace heap insert and remove max on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty heap, one item, duplicate priorities, and parent index zero.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Insert has the complexity shown in the table.

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
go test ./Heap
go test ./Heap -bench=.
```

## Next Topics

- [Sorts](../Sorts/)
- [Graph_algo](../Graph_algo/)
