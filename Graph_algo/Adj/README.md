# Graph Representation

## Goal

Compare adjacency matrix, adjacency list, and hash-map representations.

## Prerequisites

Slices, maps, and file parsing.

## Core Invariant

The vertex and edge counts must match the stored adjacency data. Undirected edges must be visible from both endpoints.

## Complexity

Matrix edge lookup: Time O(1), Space O(V^2). List or hash adjacency iteration: Time O(degree), Space O(V + E).

## Practice Tasks

- Add tests for self-loop rejection.
- Add tests for parallel-edge rejection.
- Compare memory tradeoffs for sparse and dense graphs.

## Test Command

```bash
go test ./Graph_algo/Adj
```

## Related Topics

- [Graph_algo/BFS](../BFS/)
- [Graph_algo/DFS](../DFS/)
