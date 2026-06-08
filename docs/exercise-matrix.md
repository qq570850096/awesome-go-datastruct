# Exercise Matrix

Use this matrix to move from reading code to changing code safely. Each task includes a target path and a validation command.

For the core learning spine, start from the module README first. Those textbook chapters include local exercises, hints, and reference answers, while this matrix helps you choose the next implementation or testing task.

## Stage 1: Beginner

| Task type | Task | Target | Completion standard | Validation |
| --- | --- | --- | --- | --- |
| Reading | Explain zero values, constants, and control flow. | [BasicGo](../BasicGo/) and `BasicGo/basics` | You can describe each test case in plain English. | `go test ./BasicGo/basics` |
| Rewrite | Convert one value update into a pointer update. | [BasicGo](../BasicGo/) and `BasicGo/pointers` | Tests show the caller-observed value changes. | `go test ./BasicGo/pointers` |
| Test | Add a table-driven test for a calculator edge case. | [BasicGo](../BasicGo/) and `BasicGo/testingdemo` | The new test fails before the fix and passes after it. | `go test ./BasicGo/testingdemo` |
| Reading | Compare value and pointer receivers. | [BasicGo](../BasicGo/) and `BasicGo/structs` | You can explain which receiver mutates state. | `go test ./BasicGo/structs` |

## Stage 2: Intermediate

| Task type | Task | Target | Completion standard | Validation |
| --- | --- | --- | --- | --- |
| Reading | Trace goroutine lifetime and synchronization. | `BasicGo/GoRoutine` | You can identify when work starts and completes. | `go test ./BasicGo/GoRoutine` |
| Implement | Add a channel fan-in example with cancellation. | `BasicGo/channelselect` | Closed inputs and canceled context both exit cleanly. | `go test ./BasicGo/channelselect` |
| Test | Add a context timeout test. | `BasicGo/context` | Timeout behavior is deterministic. | `go test ./BasicGo/context` |
| Race check | Run the race detector on shared state examples. | `BasicGo/sharedvars` | You understand each synchronized path. | `go test -race ./BasicGo/sharedvars` |

## Stage 3: Advanced

| Task type | Task | Target | Completion standard | Validation |
| --- | --- | --- | --- | --- |
| Test | Add empty and one-node tests. | [Linked](../Linked/) | Boundary operations do not panic unexpectedly. | `go test ./Linked` |
| Implement | Add an eviction-order scenario. | [DoubleLinked](../DoubleLinked/) | LRU/LFU/FIFO behavior is visible in tests. | `go test ./DoubleLinked` |
| Test | Verify heap order after repeated insert/remove. | [Heap](../Heap/) | Removed values follow priority order. | `go test ./Heap` |
| Refactor | Reduce duplication in tree traversal tests. | [BinarySearch](../BinarySearch/), [AVL](../AVL/), [Red-Black](../Red-Black/) | Test intent stays clear. | `go test ./BinarySearch ./AVL ./Red-Black` |
| Implement | Add graph path reconstruction coverage. | [Graph_algo/search](../Graph_algo/search/) | Paths are valid and deterministic. | `go test ./Graph_algo/search` |
| Benchmark | Add a sorting benchmark. | [Sorts](../Sorts/) | Benchmark can run without changing correctness tests. | `go test ./Sorts -bench=.` |

## Stage 4: Expert

| Task type | Task | Target | Completion standard | Validation |
| --- | --- | --- | --- | --- |
| Refactor | Improve package boundaries for a demo. | `webdemo/minigin` | Public APIs remain small and tests pass. | `go test ./webdemo/minigin` |
| Test | Add middleware behavior coverage. | `webdemo/minigin` | Middleware order and failure paths are asserted. | `go test ./webdemo/minigin` |
| Benchmark | Add heap or graph traversal benchmarks. | [Heap](../Heap/), [Graph_algo](../Graph_algo/) | Results are reproducible enough for comparison. | `go test ./Heap -bench=.` |
| Tooling | Add a documented race-detector example. | `BasicGo/sharedvars` | The command and expected lesson are documented. | `go test -race ./BasicGo/sharedvars` |
| Documentation | Add a new topic using the contribution template. | Any module | README, tests, directory, catalog, and graph are updated. | `go vet ./... && go test ./...` |
