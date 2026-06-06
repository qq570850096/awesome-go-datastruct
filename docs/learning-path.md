# Learning Path

The project can be studied in four passes. Each pass should end with `go test ./...`.

## Pass 1: Go Mechanics

Start with `BasicGo/`.

Focus:

- Basic syntax, variables, functions, control flow, slices, maps, structs, methods, and interfaces.
- Errors, `defer`, `panic`, and `recover`.
- Goroutines, channels, `select`, `sync`, and `context`.
- Generics and reflection.

Practice:

```bash
go test ./BasicGo/...
```

## Pass 2: Core Data Structures

Study the structures that teach pointer manipulation and invariants.

Order:

1. `Linked/`
2. `DoubleLinked/`
3. `stack/`
4. `queue/`
5. `Heap/`
6. `Union/`

Key questions:

- Which operation changes pointers?
- Which operation changes size?
- What does an empty structure look like?
- Which edge case should never panic?

## Pass 3: Ordered and Search Structures

Move from simple invariants to balanced or indexed structures.

Order:

1. `BinarySearch/`
2. `AVL/`
3. `Red-Black/`
4. `Segment/`
5. `Trie/`
6. `skiplists/`

Key questions:

- What invariant keeps lookup efficient?
- What happens after insertion or deletion?
- How do tests prove ordering, balance, or prefix behavior?

## Pass 4: Algorithms and Engineering

Study algorithm families and small systems.

Order:

1. `Sorts/`
2. `Graph_algo/Adj/`
3. `Graph_algo/BFS/`
4. `Graph_algo/DFS/`
5. `Graph_algo/search/`
6. `Graph_algo/leetcode/`
7. `DesignPatterns/`
8. `OSExam/`
9. `webdemo/`

Key questions:

- What input shape is worst-case?
- Which state is stored in a queue, stack, map, or recursion frame?
- Which behavior belongs in tests instead of console output?

## Weekly Plan

| Week | Goal | Packages |
| --- | --- | --- |
| 1 | Go syntax and testing | `BasicGo/`, `Utils/` |
| 2 | Pointers and linear structures | `Linked/`, `DoubleLinked/`, `stack/`, `queue/` |
| 3 | Priority and connectivity | `Heap/`, `Union/`, `Set/` |
| 4 | Trees | `BinarySearch/`, `AVL/`, `Red-Black/`, `Segment/` |
| 5 | Search structures | `Trie/`, `skiplists/` |
| 6 | Sorting | `Sorts/` |
| 7 | Graphs | `Graph_algo/` |
| 8 | Engineering topics | `DesignPatterns/`, `OSExam/`, `webdemo/` |

## How to Add Exercises

For each new exercise, add:

- One implementation file.
- One `_test.go` file with table-driven cases.
- One README section explaining the invariant and complexity.
- A note in [catalog.md](catalog.md) if it is a new topic.
