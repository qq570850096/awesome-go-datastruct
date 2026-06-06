# Awesome Go Data Structures

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Hands-on Go data structures, algorithms, language basics, design patterns, and small engineering demos. The repository is organized as a learning notebook: read the concept, run the tests, then modify the implementation until the behavior is yours.

Chinese notes are available in [README.zh.md](README.zh.md). A fuller route is in [docs/learning-path.md](docs/learning-path.md).

## Quick Start

Requires Go 1.22 or newer.

```bash
go test ./...
```

Run one topic while learning:

```bash
go test ./Linked
go test ./DoubleLinked -run TestFIFO
go test ./Graph_algo/leetcode
```

## Learning Tracks

| Track | Start here | Goal |
| --- | --- | --- |
| Go basics | `BasicGo/` | Understand syntax, testing, errors, interfaces, generics, goroutines, channels, context, and reflection. |
| Linear structures | `Linked/`, `DoubleLinked/`, `stack/`, `queue/` | Implement linked lists, stacks, queues, and cache policies such as LRU, LFU, and FIFO. |
| Trees and sets | `BinarySearch/`, `AVL/`, `Red-Black/`, `Segment/`, `Trie/`, `Union/`, `Heap/`, `skiplists/` | Compare ordered lookup structures and understand their complexity tradeoffs. |
| Sorting and graph algorithms | `Sorts/`, `Graph_algo/` | Practice classic sorting, graph representation, BFS, DFS, connectivity, cycle detection, and grid search problems. |
| Engineering topics | `DesignPatterns/`, `OSExam/`, `webdemo/` | Connect algorithmic thinking to design patterns, OS scheduling, and small web services. |

## Repository Map

| Path | What it teaches |
| --- | --- |
| `BasicGo/` | Go language features and concurrency patterns. |
| `Linked/` | Singly linked list operations and interview-style list problems. |
| `DoubleLinked/` | Doubly linked list primitives plus LRU, LFU, and FIFO cache strategies. |
| `stack/`, `queue/`, `main/622.go` | Stack, queue, and circular queue APIs. |
| `BinarySearch/`, `AVL/`, `Red-Black/`, `Segment/` | Binary search tree variants and range query structures. |
| `Trie/`, `skiplists/`, `Set/`, `Union/`, `Heap/` | Search, set, union-find, heap, and priority queue structures. |
| `Sorts/` | Bubble, insertion, selection, merge, quick, shell, and bucket/counting sort examples. |
| `Graph_algo/` | Graph storage, BFS/DFS traversal, path search, cycle detection, bipartition, and LeetCode-style graph problems. |
| `DesignPatterns/` | Creational, structural, behavioral, and compound design pattern examples. |
| `OSExam/` | FCFS, SJF, priority scheduling, and file-system skeleton exercises. |
| `webdemo/` | Small HTTP, Gin, and mini-Gin practice projects. |

For a sortable topic index, see [docs/catalog.md](docs/catalog.md).

## Project Standards

- Learning path: [docs/learning-path.md](docs/learning-path.md)
- Topic catalog: [docs/catalog.md](docs/catalog.md)
- Quality bar: [docs/project-standards.md](docs/project-standards.md)
- Reference projects: [docs/reference-projects.md](docs/reference-projects.md)
- Roadmap: [ROADMAP.md](ROADMAP.md)

## Study Loop

1. Read the topic README or [docs/catalog.md](docs/catalog.md).
2. Run only that package's tests.
3. Change one behavior or add one edge case.
4. Run `go test ./...`.
5. Write down the invariant you learned.

This mirrors strong learning repositories: small runnable examples, ordered exercises, local explanations, and a single repeatable quality gate.

## Quality Gate

Every change should keep this command green:

```bash
go vet ./...
go test ./...
```

New topics should include:

- A concise README or note explaining the invariant.
- Table-driven tests for normal, edge, and empty input cases.
- A short complexity note when the algorithm has meaningful time/space tradeoffs.

See [CONTRIBUTING.md](CONTRIBUTING.md) for the expected topic template.

## License

MIT. See [LICENSE](LICENSE).
