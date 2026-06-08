# Learning Catalog

This catalog is a learning-oriented index. `DIRECTORY.md` is the fast repository index; this file adds difficulty, prerequisites, next topics, and practice tasks.

## Textbook Chapter Spine

Use these module chapters as the primary study route. Every chapter includes worked examples, practice exercises, hints, reference answers, and test commands.

| Stage | Chapters | What the learner practices |
| --- | --- | --- |
| Foundation | [BasicGo](../BasicGo/) | Go values, pointers, slices, maps, methods, interfaces, errors, tests, and concurrency primitives. |
| Linear structures | [Linked](../Linked/), [DoubleLinked](../DoubleLinked/), [stack](../stack/), [queue](../queue/), [Circular queue](../main/) | Pointer updates, boundary cases, cache ordering, LIFO/FIFO invariants, and ring-buffer indexing. |
| Priority and connectivity | [Heap](../Heap/), [Union](../Union/) | Array-backed trees, priority order, roots, path compression, and connectivity queries. |
| Trees and search structures | [BinarySearch](../BinarySearch/), [AVL](../AVL/), [Red-Black](../Red-Black/), [Segment](../Segment/), [Trie](../Trie/), [skiplists](../skiplists/) | Ordering, balancing, rotations, range aggregation, prefix search, and randomized indexing. |
| Sorting and graphs | [Sorts](../Sorts/), [Graph_algo](../Graph_algo/), [Adj](../Graph_algo/Adj/), [BFS](../Graph_algo/BFS/), [DFS](../Graph_algo/DFS/), [search](../Graph_algo/search/), [leetcode](../Graph_algo/leetcode/) | Algorithm selection, stability, representation tradeoffs, traversal state, paths, cycles, and applied graph modeling. |

| Topic | Path | Difficulty | Prerequisites | Core idea | Next topics | Practice tasks |
| --- | --- | --- | --- | --- | --- | --- |
| Basic Go | [BasicGo](../BasicGo/) | Beginner | None | Learn syntax and runtime behavior through small packages. | Pointers, structs, tests | Add one table-driven test to a basic function. |
| Pointers | [BasicGo/pointers](../BasicGo/pointers/) | Beginner | Basic syntax | Understand address-based mutation. | Linked lists, trees | Rewrite a value-based update to a pointer-based update. |
| Structs and methods | [BasicGo/structs](../BasicGo/structs/) | Beginner | Pointers | Model data and behavior with receivers. | Interfaces, design patterns | Add a method with both value and pointer receiver tests. |
| Interfaces | [BasicGo/interface](../BasicGo/interface/) | Beginner | Structs | Express behavior through implicit contracts. | Errors, patterns | Create a small interface and two implementations. |
| Concurrency | [BasicGo](../BasicGo/) | Intermediate | Functions and tests | Learn goroutines, channels, context, and shared state. | Web demos, race checks | Run `go test -race ./BasicGo/sharedvars`. |
| Singly linked list | [Linked](../Linked/) | Advanced | Pointers | Maintain `next` links and list length. | Double linked list, caches | Add tests for empty, one-node, and duplicate cases. |
| Doubly linked list and caches | [DoubleLinked](../DoubleLinked/) | Advanced | Linked lists, maps | Maintain `prev` and `next` while enforcing eviction policy. | LRU/LFU tuning | Add an eviction-order test. |
| Stack and queue | [stack](../stack/), [queue](../queue/) | Beginner | Slices | Preserve LIFO or FIFO order. | BFS, circular queue | Add boundary tests for empty pop/dequeue. |
| Heap | [Heap](../Heap/) | Advanced | Arrays and trees | Preserve parent-child priority order. | Priority queues, Dijkstra | Add a benchmark for heap insert/remove. |
| Search trees | [BinarySearch](../BinarySearch/), [AVL](../AVL/), [Red-Black](../Red-Black/) | Advanced | Recursion | Preserve ordering and balance. | Segment tree, maps | Add rotation or deletion edge-case tests. |
| Segment tree | [Segment](../Segment/) | Advanced | Recursion, arrays | Combine child intervals into parent intervals. | Range queries | Add query tests for left-only and right-only intervals. |
| Trie | [Trie](../Trie/) | Advanced | Strings, maps | Share prefixes for lookup. | Word search | Add tests for prefix-only and missing-word cases. |
| SkipList | [skiplists](../skiplists/) | Advanced | Linked lists, probability | Use random levels for average logarithmic lookup. | Ordered maps | Add tests with deterministic random seeds where possible. |
| Union-find | [Union](../Union/) | Advanced | Arrays | Track connected components with roots and compression. | Graph connectivity | Add path compression assertions. |
| Sorting | [Sorts](../Sorts/) | Advanced | Arrays, recursion | Compare algorithmic tradeoffs and stability. | Benchmarks | Add tests for duplicates and already-sorted input. |
| Graph representation | [Graph_algo/Adj](../Graph_algo/Adj/) | Advanced | Slices, maps | Choose representation by query and memory needs. | BFS, DFS | Add tests for self-loop and parallel-edge rejection. |
| Graph traversal | [Graph_algo/BFS](../Graph_algo/BFS/), [Graph_algo/DFS](../Graph_algo/DFS/) | Advanced | Queue, recursion | Track visited state and traversal order. | Graph search | Add disconnected-graph tests. |
| Graph search | [Graph_algo/search](../Graph_algo/search/) | Advanced | BFS, DFS | Solve paths, cycles, and bipartite checks. | Applied graph problems | Add path reconstruction tests. |
| Design patterns | [DesignPatterns](../DesignPatterns/) | Expert | Interfaces, composition | Learn responsibility boundaries. | API design | Refactor one example to reduce coupling. |
| Web demos | [webdemo](../webdemo/) | Expert | HTTP, context | Understand routing and middleware. | API design, testing | Add a middleware test. |
| OS exercises | [OSExam](../OSExam/) | Expert | Sorting and state | Model scheduling and resource rules. | Systems thinking | Add a scheduling edge case. |

## Catalog Maintenance

When adding a topic, update this file, `DIRECTORY.md`, and `docs/data/knowledge-graph.json` in the same change.
