# Red-Black Tree

A red-black tree trades stricter AVL balance for cheaper updates while still guaranteeing logarithmic height.

## Learning Goals

- Understand color invariants.
- Trace rotation and color-flip cases.
- Compare red-black balance with AVL balance.

## Prerequisites

- BST ordering and rotations.

## Mental Model

Red links model temporary grouped nodes. Insertions create red links first, then rotations and color flips restore the representation.

## Diagram

```text
root is black
red links do not chain
every root-to-leaf path has equal black height
```

## Terminology

| Term | Meaning |
| --- | --- |
| Red link | A link used to encode a temporary grouped node. |
| Black height | Number of black nodes on a root-to-leaf path. |
| Color flip | A local color change that splits a temporary group. |
| Left-leaning | A convention that red links lean left. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Tree.go | Insertion, rotations, color flips, and lookup. |
| Tree_test.go | Large-scale comparison tests. |

## Core Invariants

- The root is black.
- No two red links appear consecutively.
- Every root-to-leaf path has the same black height.
- Inorder order remains sorted.

## Operation Walkthrough

Insert as in a BST, color the new node red, then repair local patterns: rotate left for right-leaning red links, rotate right for two left red links, and color-flip when both children are red.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Search | O(log n) | O(1) | Color invariants bound height. |
| Insert | O(log n) | O(1) | Local repairs along one path. |
| Rotate/color flip | O(1) | O(1) | Only local links and colors change. |

## Common Mistakes And Edge Cases

- Forgetting to force the root black.
- Breaking BST order during rotation.
- Testing speed without checking ordering.

## Worked Example

When both children of a node are red, a color flip pushes red upward. This resembles splitting a temporary 4-node in a 2-3-4 tree.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace insert repair cases on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty tree, repeated inserts, and color-flip paths.

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
go test ./Red-Black
```

## Next Topics

- [AVL](../AVL/)
- [Segment](../Segment/)
