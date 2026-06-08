# Applied Graph Exercises

This chapter collects graph and state-search exercises. The purpose is not to memorize problem numbers, but to practice turning a story into vertices, edges, states, and transitions.

## Learning Goals

- Translate problem statements into graph state.
- Choose BFS, DFS, or union-find based on the question.
- Write tests around boundary grids and unreachable states.

## Prerequisites

- Graph representation, BFS, DFS, and search state.

## Mental Model

Many grid and puzzle problems are graph problems without explicit graph structs. A cell, word, lock state, or water amount can be a vertex; legal moves are edges.

## Diagram

```text
problem state -> vertex
legal move    -> edge
visited set   -> prevents repeated work
queue/stack   -> controls exploration order
```

## Terminology

| Term | Meaning |
| --- | --- |
| State | A value that represents one position in the search space. |
| Transition | A legal move from one state to another. |
| Grid neighbor | A cell reachable by moving up, down, left, or right. |
| Visited set | Records states already processed or scheduled. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| 1091.go, 200.go, 695.go | Grid traversal examples. |
| 752.go, 4LWater.go | State-space BFS examples. |
| 785.go | Bipartite graph exercise. |
| *_test.go | Applied behavior checks. |

## Core Invariants

- Every generated neighbor must be valid for the problem.
- Visited state is recorded before repeated work explodes.
- BFS levels correspond to shortest move counts when all moves have equal cost.

## Operation Walkthrough

Start by defining the state type. Then write a function that generates valid next states. Choose BFS when the problem asks for the fewest moves, DFS when it asks for reachability or connected size, and union-find when it asks for connectivity after many unions.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Grid BFS/DFS | O(rows * cols) | O(rows * cols) | Each cell is processed at most once. |
| State-space BFS | O(states + transitions) | O(states) | Visited set bounds exploration. |
| Bipartite check | O(V + E) | O(V) | Coloring traversal. |

## Common Mistakes And Edge Cases

- Marking grid cells after enqueueing duplicates.
- Mixing row and column bounds.
- Using DFS for shortest-move problems.
- Forgetting impossible or empty input cases.

## Worked Example

In Number of Islands, each land cell is a vertex and four-direction land neighbors are edges. DFS or BFS from one unvisited land cell marks exactly one island.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace grid BFS or DFS on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty grid, one cell, all water, all land, and unreachable state.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why state expansion has the complexity shown in the table.

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
go test ./Graph_algo/leetcode
```

## Next Topics

- [Graph_algo/search](../search/)
- [docs/exercise-matrix.md](../../docs/exercise-matrix.md)
