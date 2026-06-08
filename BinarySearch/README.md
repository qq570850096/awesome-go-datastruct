# Binary Search Tree

A binary search tree stores ordered data in a shape that makes search depend on height rather than total size.

## Learning Goals

- Use ordering to guide search.
- Understand traversal orders.
- Handle deletion cases without breaking ordering.

## Prerequisites

- Recursion, pointers, and comparisons.

## Mental Model

Every node splits the remaining values into smaller values on the left and larger values on the right. Search follows one branch at each level.

## Diagram

```text
        5
      /   \
     3     8
    / \   /
   2   4 7
```

## Terminology

| Term | Meaning |
| --- | --- |
| BST property | Left subtree values are smaller and right subtree values are larger. |
| Traversal | A systematic way to visit nodes. |
| Successor | The smallest value greater than a node. |
| Height | The longest path from a node to a leaf. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Tree.go | Insert, search, traversal, min/max, and deletion. |
| Tree_test.go | Search and traversal coverage. |

## Core Invariants

- Every subtree is a valid BST.
- Inorder traversal yields sorted values.
- Deletion reconnects children without losing ordered nodes.

## Operation Walkthrough

Search compares the target with the current node and chooses exactly one branch. Deletion has three cases: leaf, one child, and two children. The two-child case replaces the node with its successor so inorder order remains valid.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Search | Average O(log n), worst O(n) | O(h) | Follows one branch per level. |
| Insert | Average O(log n), worst O(n) | O(h) | Searches for a nil child position. |
| Delete | Average O(log n), worst O(n) | O(h) | Search plus restructuring. |
| Traversal | O(n) | O(h) | Visits every node once. |

## Common Mistakes And Edge Cases

- Assuming a plain BST is always balanced.
- Deleting a two-child node without preserving successor order.
- Confusing traversal orders.

## Worked Example

Inserting sorted values 1, 2, 3 produces a chain. Inserting 2, 1, 3 produces a balanced shape. Both are valid BSTs, but their heights differ.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace search and deletion on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty tree, missing key, leaf deletion, one-child deletion, and two-child deletion.

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
go test ./BinarySearch
```

## Next Topics

- [AVL](../AVL/)
- [Red-Black](../Red-Black/)
