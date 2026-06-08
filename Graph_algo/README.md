# Graph Algorithms

This module groups graph representations, traversals, search utilities, and applied graph exercises.

| Path | Focus | Command |
| --- | --- | --- |
| [Adj](Adj/) | Matrix, list, and hash-based graph representations | `go test ./Graph_algo/Adj` |
| [BFS](BFS/) | Breadth-first traversal | `go test ./Graph_algo/BFS` |
| [DFS](DFS/) | Depth-first traversal and connected components | `go test ./Graph_algo/DFS` |
| [search](search/) | Single-source paths, cycle detection, bipartite detection | `go test ./Graph_algo/search` |
| [leetcode](leetcode/) | Applied graph and search problems | `go test ./Graph_algo/leetcode` |

## Graph Data Files

| File | Meaning |
| --- | --- |
| `g.txt` | General undirected graph |
| `g2.txt` | Graph with a cycle |
| `g2_noCycle.txt` | Cycle-free variant of `g2.txt` |
| `notBip.txt` | Non-bipartite graph example |

## Complexity Summary

| Algorithm | Time | Space |
| --- | --- | --- |
| BFS | O(V + E) | O(V) |
| DFS | O(V + E) | O(V) |
| Connected components | O(V + E) | O(V) |
| Cycle detection | O(V + E) | O(V) |
| Bipartite detection | O(V + E) | O(V) |

## Study Tasks

- Compare matrix, list, and hash representations.
- Add tests for disconnected graphs.
- Add path reconstruction tests for graph search.
