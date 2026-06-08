# Depth-First Search

Depth-first search explores as far as possible before backtracking. It is useful for connected components, cycle reasoning, and structural graph exploration.

## Learning Goals

- Understand recursive traversal state.
- Count connected components.
- Compare DFS with BFS.

## Prerequisites

- Recursion or stacks and graph representation.

## Mental Model

DFS follows one path deeply, then returns to the previous branching point. The call stack or an explicit stack remembers unfinished work.

## Diagram

```text
enter v
  enter neighbor
    enter neighbor
  backtrack
exit v
```

## Terminology

| Term | Meaning |
| --- | --- |
| Backtracking | Returning to a previous vertex after finishing a branch. |
| Component | A maximal set of mutually reachable vertices. |
| Visited mark | State that prevents revisiting. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| CC.go | Connected component traversal. |
| cc_test.go | Component tests. |

## Core Invariants

- Visited vertices are not processed again.
- Every vertex in one DFS tree belongs to the same component.
- Disconnected graphs require starting DFS from each unvisited vertex.

## Operation Walkthrough

Start DFS from an unvisited vertex and mark it with the current component id. Recursively visit unvisited neighbors. When recursion returns, all vertices reachable from the start have the same component id.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| DFS traversal | O(V + E) | O(V) | Each vertex and edge is visited a bounded number of times. |
| Connected components | O(V + E) | O(V) | Runs DFS from each unvisited component root. |

## Common Mistakes And Edge Cases

- Forgetting to start a new DFS for disconnected vertices.
- Using recursion without considering depth.
- Changing component count at the wrong time.

## Worked Example

In a graph with edges 0-1 and 2-3, DFS from 0 marks component 0, and a later DFS from 2 marks component 1.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace recursive DFS calls on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for disconnected graph, cycle, and single vertex.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why DFS traversal has the complexity shown in the table.

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
go test ./Graph_algo/DFS
```

## Next Topics

- [Graph_algo/BFS](../BFS/)
- [Graph_algo/search](../search/)
