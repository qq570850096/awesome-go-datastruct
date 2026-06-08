# Expert Track

The expert track turns repository study into engineering habits. It focuses on how Go code is tested, measured, reviewed, and maintained.

## 1. Testing Strategy

Goal: every important behavior should be repeatable.

Current anchors:

- [BasicGo/testingdemo](../BasicGo/testingdemo/)
- Package-level `_test.go` files across data structures and algorithms

Practice:

```bash
go test ./...
go test ./Graph_algo/search -run TestSingleSource
```

Next improvements:

- Add more edge-case tests for trees, caches, and graph representations.
- Convert console-only demos into assertive tests where practical.

## 2. Benchmarks

Goal: connect complexity claims to measurement.

Current anchors:

- [Sorts](../Sorts/)
- [Heap](../Heap/)
- [Union](../Union/)
- [Graph_algo](../Graph_algo/)

Suggested commands:

```bash
go test ./Sorts -bench=.
go test ./Heap -bench=.
```

TODO:

- Add benchmark functions for sorting algorithms.
- Add heap insert/remove benchmarks.
- Add union-find and graph traversal benchmarks.

## 3. Race Detector

Goal: make data-race thinking a normal part of concurrent Go development.

Current anchor:

- [BasicGo/sharedvars](../BasicGo/sharedvars/)

Command:

```bash
go test -race ./BasicGo/sharedvars
```

TODO:

- Add a documented intentionally unsafe example that demonstrates race detector output.
- Add a fixed version beside it and explain the synchronization boundary.

## 4. pprof And Performance Diagnostics

Goal: know when measurement needs profiling rather than guessing.

Current status: TODO.

Future topics:

- CPU profiles for sorting or graph traversal.
- Allocation profiles for list, heap, and trie operations.
- Documentation for interpreting profiles.

## 5. Context Propagation

Goal: understand request lifetime and cancellation.

Current anchors:

- [BasicGo/context](../BasicGo/context/)
- [webdemo](../webdemo/)

Practice:

- Add a handler or service function that accepts `context.Context`.
- Add a timeout test that proves cancellation is observed.

## 6. Error Modeling

Goal: make failures inspectable and recoverable.

Current anchors:

- [BasicGo/errors](../BasicGo/errors/)
- [BasicGo/interface](../BasicGo/interface/)

Practice:

- Prefer returning errors over panics for ordinary failures.
- Use sentinel errors or typed errors when callers need branching behavior.
- Wrap errors with `%w` when preserving the cause matters.

## 7. Package Boundaries And API Design

Goal: keep packages teachable and maintainable.

Current anchors:

- [CONTRIBUTING.md](../CONTRIBUTING.md)
- [docs/project-standards.md](project-standards.md)
- [docs/catalog.md](catalog.md)

Practice:

- Keep constructors and exported methods small.
- Keep demos separate from reusable logic.
- Update docs and knowledge graph data when adding a concept.

## Expert Completion Checklist

You are ready to contribute expert-level changes when you can:

- Add tests before or beside behavior changes.
- Explain the invariant a data structure must preserve.
- Decide whether a failure should return an error or panic.
- Run `go vet`, `go test`, and race checks when relevant.
- Add benchmark or pprof TODOs where performance claims need evidence.
