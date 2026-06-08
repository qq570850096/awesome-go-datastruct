# SkipList

A SkipList is an ordered structure that uses random levels to get balanced-tree-like average behavior with linked-list mechanics.

## Learning Goals

- Understand layered forward pointers.
- Explain expected logarithmic search.
- Compare randomized balancing with tree rotations.

## Prerequisites

- Linked lists, ordering, and probability basics.

## Mental Model

Level 0 contains every element. Higher levels skip over more elements, letting search move right quickly and then drop down when it would overshoot.

## Diagram

```text
L2: 1 -------- 9
L1: 1 ---- 5 --9
L0: 1 -3 -5 -7-9
```

## Terminology

| Term | Meaning |
| --- | --- |
| Level | One horizontal layer of forward pointers. |
| Promotion | Randomly assigning a node to higher levels. |
| Forward pointer | Pointer to the next node at the same level. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| SkipLists.go | Search, insert, delete, and random-level handling. |
| SkipLists_test.go | Behavior checks. |

## Core Invariants

- Level 0 is fully sorted.
- Higher levels are sorted subsequences of lower levels.
- Search never moves right past the target.

## Operation Walkthrough

Search starts at the highest level. Move right while the next value is still smaller than the target; otherwise drop down one level. Insert uses the same path to splice the new node into every promoted level.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Search | Expected O(log n), worst O(n) | O(1) | Random levels reduce expected steps. |
| Insert | Expected O(log n) | O(log n) | Tracks update path. |
| Delete | Expected O(log n) | O(log n) | Relinks each level containing the node. |

## Common Mistakes And Edge Cases

- Assuming random behavior is untestable.
- Breaking sorted order on high levels.
- Updating level 0 but not promoted levels.

## Worked Example

Searching for 7 may skip from 1 to 5 at a higher level, then drop to level 0 to finish.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace search path across levels on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for missing value, first value, last value, and repeated inserts.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Search has the complexity shown in the table.

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
go test ./skiplists
```

## Next Topics

- [AVL](../AVL/)
- [Red-Black](../Red-Black/)
