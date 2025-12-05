# Go Data Structures (Go 1.20+)

English overview of this repository. It collects hand-written data structures and algorithm exercises in Go, plus small concurrency demos for practice and teaching.

## Quick start
- Requires Go 1.20 or newer (`module` path is `algo`).
- Run everything: `go test ./...`
- Run a single package, e.g. contexts: `cd BasicGo/context && go test -run TestRequestIDPropagation`
- Prefer reading code alongside tests; most directories include small examples.

## What’s inside
- `BasicGo/`: Go language features and concurrency patterns (`select`, `context`, `defer/recover`, generics).
- `Linked/`: singly linked list basics and common interview tasks (reverse, deduplicate, k-th from end, cycle detection).
- `DoubleLinked/`: doubly linked list primitives plus cache strategies (LRU/LFU/FIFO).
- `stack/`, `queue/`, `main/622.go`: stacks and circular queues, including the LeetCode 622 API.
- Trees: `BinarySearch` (BST), `AVL`, `Red-Black`, `Segment`.
- `Trie/`: prefix tree for autocomplete/prefix queries.
- `Heap/`: binary heap and priority queue helpers.
- `Union/`: disjoint-set/union-find with path compression.
- `Sorts/`: bubble, insertion, selection, merge, quick, shell, bucket sort.
- `Graph_algo/`: adjacency-list graphs with BFS/DFS, cycle and bipartite detection, plus leetcode-style problems.
- Utilities: `Utils/Interfaces` for type-safe comparisons; `pkg/` keeps a minimal module cache layout.

## Notes
- The original Chinese guide is now at `README.zh.md`.
- License: MIT (see `LICENSE`).
