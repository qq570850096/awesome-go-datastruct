# Beginner to Expert Go Roadmap

This roadmap treats the repository as a runnable workshop. Each stage asks you to read concepts, inspect code, run tests, and make small changes until the underlying invariant or engineering boundary becomes clear.

## How To Use This Roadmap

1. Identify the stage that matches your current skill level.
2. For each topic, read the local README and run the tests.
3. Make one small change, then restore confidence with `go test`.
4. Before moving forward, explain the key invariant or tradeoff in your own words.

Repository-wide quality gate:

```bash
go vet ./...
go test ./...
```

## Stage 1: Go Language Foundations

Goal: build fluency with syntax, types, functions, pointers, structs, interfaces, errors, and tests.

| Topic | Repository path | Learning outcome | Validation command |
| --- | --- | --- | --- |
| Syntax and control flow | [BasicGo/basics](../../BasicGo/basics/) | Understand variables, constants, loops, branches, functions, and package-level code. | `go test ./BasicGo/basics` |
| Pointers | [BasicGo/pointers](../../BasicGo/pointers/) | Understand value semantics, address semantics, and pointer updates. | `go test ./BasicGo/pointers` |
| Slices and maps | [BasicGo/slicemap](../../BasicGo/slicemap/) | Understand reference-like behavior, growth, traversal, and common map patterns. | `go test ./BasicGo/slicemap` |
| Structs and methods | [BasicGo/structs](../../BasicGo/structs/) | Understand fields, receivers, embedding, and tags. | `go test ./BasicGo/structs` |
| Interfaces | [BasicGo/interface](../../BasicGo/interface/) | Understand implicit implementation, dynamic dispatch, and the `error` interface. | `go test ./BasicGo/interface` |
| Errors and defer | [BasicGo/errors](../../BasicGo/errors/) and [BasicGo/defer](../../BasicGo/defer/) | Understand wrapping, cleanup, panic boundaries, and recovery. | `go test ./BasicGo/errors ./BasicGo/defer` |
| Testing | [BasicGo/testingdemo](../../BasicGo/testingdemo/) | Write table-driven tests and run one package at a time. | `go test ./BasicGo/testingdemo` |

Completion signals:

- You can explain value passing versus pointer passing.
- You can add a table-driven test for a small function.
- You know when a pointer receiver is required.
- You can choose between returning an error and using `panic`.

## Stage 2: Concurrency And Engineering Basics

Goal: understand goroutines, channels, `select`, `context`, shared variables, and race detection as engineering tools.

| Topic | Repository path | Learning outcome | Validation command |
| --- | --- | --- | --- |
| Goroutines | [BasicGo/GoRoutine](../../BasicGo/GoRoutine/) | Understand lightweight concurrent execution and waiting. | `go test ./BasicGo/GoRoutine` |
| Channels and select | [BasicGo/channelselect](../../BasicGo/channelselect/) | Understand communication, timeouts, fan-in, and close semantics. | `go test ./BasicGo/channelselect` |
| Context | [BasicGo/context](../../BasicGo/context/) | Understand cancellation, timeout propagation, and request scope. | `go test ./BasicGo/context` |
| Shared variables | [BasicGo/sharedvars](../../BasicGo/sharedvars/) | Understand locks, atomics, `sync.Once`, and data races. | `go test ./BasicGo/sharedvars` |
| Reflection and generics | [BasicGo/reflect](../../BasicGo/reflect/) and [BasicGo/generics](../../BasicGo/generics/) | Understand runtime type inspection and type parameters. | `go test ./BasicGo/reflect ./BasicGo/generics` |
| Low-level layout | [BasicGo/lowlevel](../../BasicGo/lowlevel/) | Understand size, alignment, offsets, and `unsafe` boundaries. | `go test ./BasicGo/lowlevel` |

Useful extra check:

```bash
go test -race ./BasicGo/sharedvars
```

Completion signals:

- You can draw the data flow of a concurrent program.
- You can decide between channel communication and shared-state synchronization.
- You reach for `context` when work needs cancellation or timeouts.
- You understand what the race detector can and cannot prove.

## Stage 3: Data Structures And Algorithms

Goal: master invariants, complexity, and test strategies for common data structures and algorithms.

| Topic | Repository path | Core invariant | Validation command |
| --- | --- | --- | --- |
| Singly linked list | [Linked](../../Linked/) | Each node points to the next node and empty-list behavior is explicit. | `go test ./Linked` |
| Doubly linked list and caches | [DoubleLinked](../../DoubleLinked/) | `prev` and `next` are updated together; eviction order is testable. | `go test ./DoubleLinked` |
| Stack and queue | [stack](../../stack/) and [queue](../../queue/) | LIFO and FIFO ordering are never violated. | `go test ./stack ./queue` |
| Heap | [Heap](../../Heap/) | Parent and child priorities match the heap rule. | `go test ./Heap` |
| Search trees | [BinarySearch](../../BinarySearch/), [AVL](../../AVL/), [Red-Black](../../Red-Black/) | Ordering and balancing keep lookup efficient. | `go test ./BinarySearch ./AVL ./Red-Black` |
| Segment tree | [Segment](../../Segment/) | Parent intervals are derived from child intervals. | `go test ./Segment` |
| Trie and SkipList | [Trie](../../Trie/) and [skiplists](../../skiplists/) | Prefix sharing and layered indexing reduce lookup cost. | `go test ./Trie ./skiplists` |
| Union-find | [Union](../../Union/) | Representative roots and path compression preserve connectivity. | `go test ./Union` |
| Sorting | [Sorts](../../Sorts/) | Output is ordered and contains the same elements as input. | `go test ./Sorts` |
| Graphs | [Graph_algo](../../Graph_algo/) | Representation, visited state, paths, and connectivity agree. | `go test ./Graph_algo/...` |

Completion signals:

- You can test empty, single-element, duplicate, missing, and capacity-boundary cases.
- You can explain time and space complexity for important operations.
- You can identify the exact statements that preserve an invariant.
- You can compare tradeoffs such as heap versus sorting or BFS versus DFS.

## Stage 4: Expert Engineering Track

Goal: move from writing Go code to maintaining Go systems.

| Topic | Repository path | Training focus | Validation |
| --- | --- | --- | --- |
| Testing strategy | [BasicGo/testingdemo](../../BasicGo/testingdemo/) and package tests | Turn behavior into repeatable tests. | `go test ./...` |
| Benchmarks | [Sorts](../../Sorts/), [Heap](../../Heap/), [Graph_algo](../../Graph_algo/) | Add benchmark coverage for performance-sensitive code. | TODO in [Expert Track](../expert-track.md) |
| Vet and race detection | [BasicGo/sharedvars](../../BasicGo/sharedvars/) | Treat static checks and race detection as local gates. | `go vet ./...`, `go test -race ./BasicGo/sharedvars` |
| Design patterns | [DesignPatterns](../../DesignPatterns/) | Understand responsibility boundaries, not just pattern names. | `go test ./DesignPatterns/...` |
| Web demos | [webdemo](../../webdemo/) | Understand HTTP, routing, middleware, and framework boundaries. | `go test ./webdemo/...` |
| OS exercises | [OSExam](../../OSExam/) | Express scheduling and resource-management rules with tests. | `go test ./OSExam` |
| Package boundaries | [CONTRIBUTING.md](../../CONTRIBUTING.md) | Add docs, tests, directory entries, and graph data for new topics. | Manual review |

Completion signals:

- You can decide which package should own a new feature.
- You can design a topic README, tests, complexity notes, and exercises.
- You run `go vet`, `go test`, and `go test -race` by default.
- You know when benchmark or pprof work is needed and can write it as a clear TODO.

## Suggested 8-Week Plan

| Week | Goal | Recommended paths |
| --- | --- | --- |
| 1 | Go basics and testing | `BasicGo/basics`, `BasicGo/testingdemo` |
| 2 | Pointers, structs, interfaces, and errors | `BasicGo/pointers`, `BasicGo/structs`, `BasicGo/interface`, `BasicGo/errors` |
| 3 | Concurrency and engineering basics | `BasicGo/GoRoutine`, `BasicGo/channelselect`, `BasicGo/context`, `BasicGo/sharedvars` |
| 4 | Linear structures and caches | `Linked`, `DoubleLinked`, `stack`, `queue` |
| 5 | Trees, heap, and union-find | `Heap`, `BinarySearch`, `AVL`, `Red-Black`, `Union` |
| 6 | Sorting, Trie, SkipList, and segment tree | `Sorts`, `Trie`, `skiplists`, `Segment` |
| 7 | Graph algorithms | `Graph_algo/Adj`, `Graph_algo/BFS`, `Graph_algo/DFS`, `Graph_algo/search` |
| 8 | Engineering practice | `DesignPatterns`, `webdemo`, `OSExam`, repository-wide tests |

## Deepening Loop

For every topic:

1. Read the README and source code.
2. Run the package tests.
3. Identify the core invariant.
4. Add one edge-case test.
5. Write down the complexity.
6. Pick a next task from the [Exercise Matrix](../exercise-matrix.md).
