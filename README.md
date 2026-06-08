# Awesome Go Data Structures

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> A hands-on Go learning repository for language fundamentals, data structures, algorithms, concurrency, testing, and engineering practice.

This repository is not just a collection of code snippets. It is a runnable learning path that helps readers move from Go beginner to Go engineering expert. The route starts with language mechanics, continues through data structures and algorithms, and then moves into concurrency, testing, design patterns, Web demos, tooling, and performance awareness.

## Project Positioning

- Build a practical path from Go fundamentals to expert-level engineering habits.
- Use data structures and algorithms as the training ground for pointers, invariants, tests, and complexity analysis.
- Borrow navigation ideas from mature learning repositories such as TheAlgorithms/Python, while keeping a Go engineering focus.
- Keep every major topic connected to code, tests, practice tasks, and next-step guidance.

## Who This Is For

| Reader | Suggested entry |
| --- | --- |
| New Go learner | Start with `BasicGo/` and [Beginner to Expert](docs/roadmaps/beginner-to-expert.md). |
| Developer from another language | Read the [Knowledge Graph](docs/knowledge-graph.md) and skip concepts you already know. |
| Algorithm learner | Start with `Linked/`, `Sorts/`, `Graph_algo/`, and [DIRECTORY.md](DIRECTORY.md). |
| Backend engineer | Focus on concurrency, `context`, testing, design patterns, `webdemo/`, and the [Expert Track](docs/expert-track.md). |

## 30-Second Quick Start

Requires Go 1.22 or newer.

```bash
go vet ./...
go test ./...
```

Run one topic only:

```bash
go test ./BasicGo/pointers
go test ./DoubleLinked -run TestFIFO
go test ./Graph_algo/search
```

## Four-Stage Overview

| Stage | Goal | Recommended modules | Completion signal |
| --- | --- | --- | --- |
| Stage 1 Beginner | Learn syntax, functions, pointers, structs, interfaces, errors, and tests. | `BasicGo/` | You can modify basic examples and write table-driven tests. |
| Stage 2 Intermediate | Understand goroutines, channels, `select`, `context`, shared variables, and race thinking. | `BasicGo/GoRoutine`, `BasicGo/channelselect`, `BasicGo/context`, `BasicGo/sharedvars` | You can explain data flow and synchronization boundaries. |
| Stage 3 Advanced | Master linked lists, stacks, queues, heaps, trees, graphs, sorting, union-find, Trie, and SkipList. | `Linked/`, `Heap/`, `Graph_algo/`, `Sorts/` | You can test invariants and explain complexity tradeoffs. |
| Stage 4 Expert | Build habits around tests, benchmarks, vet, race detection, modularity, API design, and Web engineering. | `DesignPatterns/`, `webdemo/`, `OSExam/` | You can extend modules using a maintainable engineering standard. |

See the full route in [docs/roadmaps/beginner-to-expert.md](docs/roadmaps/beginner-to-expert.md).

## Textbook Core Spine

The core modules now read like textbook chapters. Each chapter includes goals, prerequisites, a mental model, invariants, operation walkthroughs, complexity derivations, common mistakes, worked examples, exercises, hints, reference answers, and test commands.

| Spine area | Start here | Continue with |
| --- | --- | --- |
| Go fundamentals | [BasicGo](BasicGo/) | Pointers, structs, interfaces, errors, tests, concurrency |
| Linear structures | [Linked](Linked/) | [DoubleLinked](DoubleLinked/), [stack](stack/), [queue](queue/), [Circular queue](main/) |
| Priority and connectivity | [Heap](Heap/) | [Union](Union/) |
| Trees and search structures | [BinarySearch](BinarySearch/) | [AVL](AVL/), [Red-Black](Red-Black/), [Segment](Segment/), [Trie](Trie/), [skiplists](skiplists/) |
| Sorting | [Sorts](Sorts/) | Benchmarks and stability analysis |
| Graphs | [Graph_algo](Graph_algo/) | [Adjacency](Graph_algo/Adj/), [BFS](Graph_algo/BFS/), [DFS](Graph_algo/DFS/), [Search](Graph_algo/search/), [Applied problems](Graph_algo/leetcode/) |

Maintainers should follow the [Textbook Style Guide](docs/textbook-style-guide.md) when adding or expanding chapters.

## Learning Entrypoints

| Entrypoint | Purpose |
| --- | --- |
| [Textbook Style Guide](docs/textbook-style-guide.md) | Defines the module chapter structure, exercise format, diagrams, complexity notation, and linking rules. |
| [Knowledge Graph](docs/knowledge-graph.md) | Shows prerequisite, practice, and deepening relationships. |
| [Structured Knowledge Data](docs/data/knowledge-graph.json) | Stable data for future generated pages or visualizations. |
| [DIRECTORY.md](DIRECTORY.md) | TheAlgorithms-style topic directory for the whole repository. |
| [Learning Catalog](docs/catalog.md) | Difficulty, prerequisites, next topics, and practice tasks. |
| [Go Programming Language Gap Map](docs/go-bible-gap-map.md) | Maps book chapters to repository stages and modules. |
| [Exercise Matrix](docs/exercise-matrix.md) | Reading, rewrite, test, implementation, benchmark, and refactor tasks. |
| [Expert Track](docs/expert-track.md) | A path from writing Go code to maintaining Go systems. |
| [Contributing Guide](CONTRIBUTING.md) | Rules for new topics, tests, docs, and knowledge graph updates. |

## Repository Map

| Path | Learning content |
| --- | --- |
| `BasicGo/` | Go basics, pointers, concurrency, context, reflection, generics, and low-level layout. |
| `Linked/`, `DoubleLinked/` | Singly linked lists, doubly linked lists, LRU, LFU, and FIFO caches. |
| `stack/`, `queue/`, `main/622.go` | Stack, queue, and circular queue. |
| `BinarySearch/`, `AVL/`, `Red-Black/`, `Segment/` | Search trees, balanced trees, and segment trees. |
| `Trie/`, `skiplists/`, `Union/`, `Heap/`, `Set/` | Prefix trees, skip lists, union-find, heaps, and sets. |
| `Sorts/` | Bubble, insertion, selection, merge, quick, shell, bucket, and counting sorts. |
| `Graph_algo/` | Graph representations, BFS, DFS, paths, cycle detection, bipartite checks, and graph exercises. |
| `DesignPatterns/` | Creational, structural, behavioral, and compound design patterns. |
| `OSExam/` | FCFS, SJF, priority scheduling, and a file-system skeleton. |
| `webdemo/` | Native HTTP, Gin, mini-Gin, and a red-packet demo. |

## Recommended Learning Loop

1. Find your current position in the [Knowledge Graph](docs/knowledge-graph.md).
2. Pick a stage in [Beginner to Expert](docs/roadmaps/beginner-to-expert.md).
3. Read the local README and source code for the module.
4. Run the module tests, for example `go test ./Graph_algo/search`.
5. Complete a task from the [Exercise Matrix](docs/exercise-matrix.md).
6. Finish with `go vet ./...` and `go test ./...`.

## Quality Bar

Every change should keep these commands green:

```bash
go vet ./...
go test ./...
```

New topics should include:

- A topic README or equivalent documentation.
- Table-driven tests.
- Complexity and core invariant notes.
- Entries in `DIRECTORY.md`, `docs/catalog.md`, and `docs/data/knowledge-graph.json`.

## Roadmap

See [ROADMAP.md](ROADMAP.md). Current priorities include generated directory indexes, more benchmarks, `go test -race` examples, Dependabot vulnerability cleanup, and a stronger expert learning track.

## License

MIT. See [LICENSE](LICENSE).
