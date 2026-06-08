# Graph Algorithms

Graphs model relationships. This chapter ties together representation, traversal, search, path reconstruction, cycle detection, bipartite checks, and applied problem solving.

## Learning Goals

- Choose a graph representation for a problem.
- Understand visited state and traversal order.
- Use BFS, DFS, and union-find when appropriate.

## Prerequisites

- Queues, stacks or recursion, maps, and sets.

## Mental Model

A graph is vertices plus edges. Most graph algorithms are disciplined ways to decide what state is visited, what frontier remains, and what extra information must be recorded.

## Diagram

```text
0 -- 1 -- 2
|    |
3 -- 4
vertices: 0..4
edges: relationships
```

## Terminology

| Term | Meaning |
| --- | --- |
| Vertex | An entity in the graph. |
| Edge | A relationship between vertices. |
| Frontier | Discovered but not fully processed vertices. |
| Visited | State proving a vertex does not need to be processed again. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Adj/ | Matrix, list, and hash graph representations. |
| BFS/ | Breadth-first traversal. |
| DFS/ | Depth-first traversal and connected components. |
| search/ | Paths, cycle detection, and bipartite checks. |
| leetcode/ | Applied graph and state-search exercises. |

## Core Invariants

- Visited state prevents infinite revisits.
- Every traversed edge belongs to the graph representation.
- Path reconstruction records enough predecessor information.

## Operation Walkthrough

Pick representation first. Then choose BFS for shortest paths in unweighted graphs, DFS for deep exploration and connected components, or union-find for connectivity queries without path details.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Adjacency scan | O(degree) | O(1) | List/hash representations expose neighbors. |
| BFS/DFS | O(V + E) | O(V) | Each vertex and edge is processed a bounded number of times. |
| Path reconstruction | O(length) | O(length) | Follows predecessor links. |

## Common Mistakes And Edge Cases

- Marking visited too late and enqueuing duplicates.
- Using DFS when BFS distance is required.
- Forgetting disconnected components.
- Mixing directed and undirected assumptions.

## Worked Example

In an unweighted graph, BFS from source 0 records predecessors. To reconstruct a path to 4, follow `pre[4]`, then `pre[...]`, until source.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace BFS frontier and visited set on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for disconnected graph, self-loop, cycle, and non-bipartite graph.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why BFS or DFS traversal has the complexity shown in the table.

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
go test ./Graph_algo/...
```

## Next Topics

- [Graph_algo/Adj](Adj/)
- [Graph_algo/BFS](BFS/)
- [Graph_algo/search](search/)
