# Breadth-First Search

Breadth-first search explores a graph by distance from the source. It is the standard tool for shortest paths in unweighted graphs.

## Learning Goals

- Use a queue frontier.
- Mark visited at the correct time.
- Connect BFS order to unweighted distance.

## Prerequisites

- Queues and graph representation.

## Mental Model

BFS processes all vertices at distance d before vertices at distance d+1. The queue stores the frontier in discovery order.

## Diagram

```text
source -> distance 0
neighbors -> distance 1
next layer -> distance 2
```

## Terminology

| Term | Meaning |
| --- | --- |
| Source | The starting vertex. |
| Layer | Vertices at the same distance from source. |
| Frontier | Queue of discovered vertices waiting to be processed. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| traverse.go | BFS traversal implementation. |
| traverse_test.go | Traversal behavior tests. |

## Core Invariants

- A vertex is marked visited before or when it enters the queue.
- The queue processes vertices in nondecreasing distance.
- Every reachable vertex is eventually processed.

## Operation Walkthrough

Initialize the queue with the source. Repeatedly dequeue one vertex, inspect neighbors, mark unvisited neighbors, and enqueue them. The queue naturally preserves layer order.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Traversal | O(V + E) | O(V) | Each vertex enters the queue at most once. |
| Neighbor scan | O(E) | O(1) | All adjacency lists are scanned across the run. |

## Common Mistakes And Edge Cases

- Marking visited only after dequeue and creating duplicate queue entries.
- Expecting BFS to solve weighted shortest paths.
- Forgetting disconnected vertices.

## Worked Example

From source 0 in a line 0-1-2, BFS visits 0, then 1, then 2. The predecessor of 2 is 1.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace BFS queue changes on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for single vertex, disconnected graph, and graph with cycles.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why BFS traversal has the complexity shown in the table.

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
go test ./Graph_algo/BFS
```

## Next Topics

- [queue](../../queue/)
- [Graph_algo/search](../search/)
