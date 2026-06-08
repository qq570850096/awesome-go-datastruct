# AVL Tree

An AVL tree is a binary search tree that keeps height tightly balanced through rotations.

## Learning Goals

- Compute height and balance factor.
- Recognize LL, RR, LR, and RL cases.
- Use rotations to restore logarithmic height.

## Prerequisites

- Binary search tree insertion and recursion.

## Mental Model

AVL insertion first behaves like BST insertion. During the return from recursion, each ancestor checks whether the height difference is too large and rotates if needed.

## Diagram

```text
LL before:          after right rotate:
      30                    20
     /                     /  \
   20                    10    30
  /
10
```

## Terminology

| Term | Meaning |
| --- | --- |
| Balance factor | Height(left) - Height(right). |
| Rotation | A local pointer rearrangement preserving inorder order. |
| LL/RR | Single-rotation imbalance cases. |
| LR/RL | Double-rotation imbalance cases. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Tree.go | Height tracking, insertion, deletion, and rotations. |
| Tree_test.go | Rotation and behavior tests. |

## Core Invariants

- BST ordering still holds.
- For every node, subtree heights differ by at most one.
- Rotations preserve inorder order.

## Operation Walkthrough

After insertion, update height on the way back up. If a node becomes too left-heavy or right-heavy, inspect the child direction to choose single or double rotation.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Search | O(log n) | O(h) | Height is logarithmic. |
| Insert | O(log n) | O(h) | One search path plus constant rotations. |
| Delete | O(log n) | O(h) | May rebalance on the return path. |
| Rotation | O(1) | O(1) | Only local pointers and heights change. |

## Common Mistakes And Edge Cases

- Updating pointers but not heights.
- Using a single rotation for LR or RL.
- Testing only one rotation shape.

## Worked Example

Insert 30, 20, 10. The root 30 becomes left-heavy with a left-heavy child, so a right rotation makes 20 the root.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace LL, RR, LR, and RL rotations on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty tree, duplicate-like insertion policy, and each rotation case.

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
go test ./AVL
```

## Next Topics

- [Red-Black](../Red-Black/)
- [BinarySearch](../BinarySearch/)
