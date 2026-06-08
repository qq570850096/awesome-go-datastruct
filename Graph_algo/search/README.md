# Graph Search Problems

Graph search builds on traversal by recording extra facts: predecessors for paths, parent edges for cycle detection, and colors for bipartite checks.

## Learning Goals

- Reconstruct paths from predecessor arrays.
- Detect cycles in undirected graphs.
- Use coloring to test bipartite structure.

## Prerequisites

- BFS, DFS, arrays, and graph representation.

## Mental Model

Traversal answers what is reachable. Search problems add a small amount of state so the traversal answers a more specific question.

## Diagram

```text
pre[v] records where v came from
color[v] records one of two partitions
parent prevents treating the edge back to parent as a cycle
```

## Terminology

| Term | Meaning |
| --- | --- |
| Predecessor | Previous vertex on a discovered path. |
| Parent edge | The edge used to enter a vertex during DFS. |
| Bipartite | A graph whose vertices can be colored with two colors so every edge crosses colors. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| SingleSourcePath.go | Single-source path state. |
| Cycle.go | Cycle detection. |
| BipartitionDetection.go | Bipartite coloring. |
| search_test.go | Search behavior tests. |

## Core Invariants

- A reconstructed path follows predecessor links back to the source.
- Cycle detection ignores the immediate parent edge.
- Adjacent vertices in a bipartite graph have different colors.

## Operation Walkthrough

For paths, traversal records `pre[neighbor] = current` when the neighbor is first discovered. For cycle detection, DFS remembers the parent vertex. For bipartite checks, traversal assigns the opposite color to each unvisited neighbor.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Single-source path setup | O(V + E) | O(V) | Traversal plus predecessor array. |
| Cycle detection | O(V + E) | O(V) | DFS with visited and parent state. |
| Bipartite detection | O(V + E) | O(V) | Color array plus traversal. |

## Common Mistakes And Edge Cases

- Overwriting predecessor after first discovery.
- Reporting the parent edge as a cycle.
- Forgetting to start checks in every disconnected component.

## Worked Example

If `pre[6]=5`, `pre[5]=2`, and `pre[2]=0`, then the path from 0 to 6 is `0 -> 2 -> 5 -> 6` after reversing the collected chain.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace predecessor path reconstruction on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for unreachable target, disconnected graph, odd cycle, and even cycle.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why path reconstruction or bipartite coloring has the complexity shown in the table.

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
go test ./Graph_algo/search
```

## Next Topics

- [Graph_algo/leetcode](../leetcode/)
- [Union](../../Union/)
