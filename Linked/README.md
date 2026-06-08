# Singly Linked List

A singly linked list is the first pointer-heavy data structure in the learning spine. It teaches how local pointer assignments change a whole reachable sequence.

## Learning Goals

- Represent a sequence with nodes and `next` pointers.
- Trace insertion, deletion, reversal, and two-pointer techniques.
- Protect nil and one-node cases with tests.

## Prerequisites

- Go pointers and structs.
- Table-driven tests.

## Mental Model

A linked list is a chain. You do not move elements in memory; you change which node points to which next node.

## Diagram

```text
head -> [1|next] -> [2|next] -> [3|nil]
```

## Terminology

| Term | Meaning |
| --- | --- |
| Node | A value plus a pointer to the next node. |
| Head | The first reachable node or a dummy node before it. |
| Predecessor | The node before the node being changed. |
| Two pointers | A fast/slow or previous/current pair. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| List.go | Core list operations. |
| 1.1 Reverse.go to 1.11 CheckIntersect.go | Focused linked-list exercises. |
| List_test.go | Behavior checks. |
| linked-list-exercises.md | Exercise index. |

## Core Invariants

- The head reaches every node exactly once unless a cycle exercise is intentional.
- Insertion and deletion must not lose required nodes.
- A nil list and one-node list are valid inputs for most algorithms.

## Operation Walkthrough

For insertion, save the predecessor old next pointer, point the new node to it, then point the predecessor to the new node. For reversal, save `next`, redirect `cur.Next`, then advance `prev` and `cur`.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Search | O(n) | O(1) | May inspect every node. |
| Insert after known node | O(1) | O(1) | Only pointer assignments. |
| Delete by index | O(n) | O(1) | Must find predecessor. |
| Reverse | O(n) | O(1) | Visits each node once. |

## Common Mistakes And Edge Cases

- Overwriting `cur.Next` before saving it.
- Returning the old head after head changes.
- Forgetting nil and one-node cases.
- Creating an accidental cycle.

## Worked Example

To reverse `1 -> 2 -> 3`, start with `prev=nil` and `cur=1`. Save `next=2`, set `1.Next=nil`, then move forward. After three iterations, `prev` is the new head.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace list reversal on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for nil list, one-node list, duplicate values, and missing values.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why reversal has the complexity shown in the table.

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
go test ./Linked
```

## Next Topics

- [DoubleLinked](../DoubleLinked/)
- [stack](../stack/)
- [queue](../queue/)
