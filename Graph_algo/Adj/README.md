# Graph Representation

Graph representation determines memory cost, edge lookup cost, and traversal ergonomics. This chapter compares matrix, list, and hash-based forms.

## Learning Goals

- Compare matrix, table, and hash representations.
- Validate edge and vertex counts.
- Choose representation based on graph density.

## Prerequisites

- Slices, maps, and file parsing.

## Mental Model

A dense graph benefits from quick matrix lookup. A sparse graph usually benefits from adjacency lists or maps because it stores only existing edges.

## Diagram

```text
matrix[u][v] answers edge lookup
list[u] stores neighbors of u
hash[u] stores neighbors of u with flexible keys
```

## Terminology

| Term | Meaning |
| --- | --- |
| Dense graph | Many possible edges exist. |
| Sparse graph | Few possible edges exist. |
| Adjacency | The neighbors of one vertex. |
| Self-loop | An edge from a vertex to itself. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Matrix.go | Matrix representation. |
| Table.go | Slice table representation. |
| Hash.go | Hash-map representation. |
| *_test.go | Representation validation. |

## Core Invariants

- Vertex count is non-negative.
- Undirected edges appear from both endpoints.
- Self-loops and parallel edges follow the module policy.

## Operation Walkthrough

Read file parsing first, then inspect how each representation stores the same edge. Compare edge lookup with neighbor iteration; these are the two operations that drive representation choice.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Matrix edge lookup | O(1) | O(V^2) | Direct cell access. |
| List neighbor scan | O(degree) | O(V + E) | Stores actual neighbors. |
| Hash neighbor scan | O(degree) | O(V + E) | Map lookup finds adjacency bucket. |

## Common Mistakes And Edge Cases

- Using a matrix for a very sparse graph without considering memory.
- Forgetting to add both directions in an undirected graph.
- Ignoring invalid vertex IDs.

## Worked Example

For vertices 0, 1, 2 and edge 0-2, matrix sets two cells, while list/hash append 2 to 0 neighbors and 0 to 2 neighbors.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace adding one undirected edge on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for negative counts, invalid vertex, self-loop, and parallel edge.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why edge insertion or lookup has the complexity shown in the table.

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
go test ./Graph_algo/Adj
```

## Next Topics

- [Graph_algo/BFS](../BFS/)
- [Graph_algo/DFS](../DFS/)
