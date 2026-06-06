# Catalog

This catalog is the fast index for learners and maintainers.

| Topic | Path | Core idea | Typical complexity |
| --- | --- | --- | --- |
| Basic Go | `BasicGo/` | Language syntax and runtime behavior | Varies |
| Singly linked list | `Linked/` | `next` pointers and dummy heads | Search O(n), insert/delete by node O(1) |
| Doubly linked list | `DoubleLinked/` | `prev` and `next` pointers | Move/delete by node O(1) |
| LRU cache | `DoubleLinked/LRU.go` | Map plus doubly linked list | Get/Put O(1) |
| LFU cache | `DoubleLinked/LFU.go` | Frequency buckets plus linked lists | Get/Put O(1) average in this model |
| FIFO cache | `DoubleLinked/FIFO.go` | Evict oldest inserted item | Get/Put O(1) |
| Stack | `stack/` | Last-in first-out | Push/Pop O(1) |
| Queue | `queue/` | First-in first-out | Enqueue/Dequeue O(1) |
| Binary search tree | `BinarySearch/` | Left smaller, right larger | Average O(log n), worst O(n) |
| AVL tree | `AVL/` | Height-balanced BST | Search/insert/delete O(log n) |
| Red-black tree | `Red-Black/` | Color-balanced BST | Search/insert/delete O(log n) |
| Segment tree | `Segment/` | Store merged interval values | Query/update O(log n) |
| Trie | `Trie/` | Shared prefixes | Search O(length) |
| Skip list | `skiplists/` | Randomized layered list | Average O(log n) |
| Heap | `Heap/` | Complete tree stored in array | Add/remove O(log n) |
| Union-find | `Union/` | Parent array and path compression | Nearly O(1) amortized |
| Sorting | `Sorts/` | Compare and non-compare sorts | Varies |
| Graph representation | `Graph_algo/Adj/` | Matrix, slice table, hash table | Depends on representation |
| BFS/DFS | `Graph_algo/BFS/`, `Graph_algo/DFS/` | Queue or recursion traversal | O(V + E) |
| Graph search | `Graph_algo/search/` | Paths, cycles, bipartition | O(V + E) |
| Design patterns | `DesignPatterns/` | Reusable object collaboration shapes | Varies |
| OS scheduling | `OSExam/` | FCFS, SJF, priority scheduling | Varies |
| Web demos | `webdemo/` | HTTP, Gin, mini framework | Varies |

## Topic Checklist

Use this checklist when improving a topic:

- Does the README say what invariant matters?
- Do tests cover empty input, one element, duplicates, and normal cases?
- Can a learner run only that package?
- Is console output limited to demos rather than core tests?
- Are artifacts and generated files excluded from git?
