# Union-Find

Union-find models connected components with parent links. It is compact, fast, and central to graph connectivity problems.

## Learning Goals

- Represent sets with parent arrays.
- Use path compression to flatten trees.
- Apply union-find to connectivity.

## Prerequisites

- Arrays and tree-shaped parent links.

## Mental Model

Each element points toward a representative root. Union connects two roots; find follows parents until it reaches a root.

## Diagram

```text
parent: [0,0,2,2]
0 and 1 share root 0
2 and 3 share root 2
union(1,2) connects both components
```

## Terminology

| Term | Meaning |
| --- | --- |
| Root | An element whose parent is itself. |
| Find | Return the representative root. |
| Union | Merge two components. |
| Path compression | Make visited nodes point closer to the root. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Find.go | Union-find implementation. |
| Find_test.go | Connectivity and performance-oriented checks. |

## Core Invariants

- Connected elements have the same root.
- Path compression must not change connectivity.
- Union should connect roots, not arbitrary intermediate nodes.

## Operation Walkthrough

Read find first. It climbs parent links until a root appears. Then read union: it finds both roots and connects one component to the other. Finally, inspect path compression as a performance repair that preserves answers.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Find | Nearly O(1) amortized | O(1) | Compression and weighting make practical cost tiny. |
| Union | Nearly O(1) amortized | O(1) | Two finds plus one parent update. |
| Connected | Nearly O(1) amortized | O(1) | Two finds and comparison. |

## Common Mistakes And Edge Cases

- Comparing immediate parents instead of roots.
- Compressing before the root is known.
- Forgetting invalid index behavior.

## Worked Example

If `0-1` and `2-3` are connected separately, `union(1,2)` makes all four items share one root.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace find with path compression on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for single element, already-connected elements, and chained parents.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Find has the complexity shown in the table.

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
go test ./Union
```

## Next Topics

- [Graph_algo/search](../Graph_algo/search/)
- [Graph_algo/leetcode](../Graph_algo/leetcode/)
