# The Go Programming Language Gap Map

This document maps the chapter structure of *The Go Programming Language* to the repository roadmap and knowledge graph. It is not a replacement for the book. It shows where each book topic can be practiced with runnable repository code.

References:

- Book site: <https://www.gopl.io/>
- Learning route: [Beginner to Expert](roadmaps/beginner-to-expert.md)
- Knowledge data: [knowledge-graph.json](data/knowledge-graph.json)

## Chapter Mapping

| Book chapter or area | Roadmap stage | Knowledge node | Repository coverage | Status |
| --- | --- | --- | --- | --- |
| Chapter 1 Tutorial | Stage 1 | `go.basics`, `go.testing` | [BasicGo/basics](../BasicGo/basics/), [BasicGo/testingdemo](../BasicGo/testingdemo/) | Covered |
| Chapter 2 Program Structure | Stage 1 | `go.basics` | [BasicGo/basics](../BasicGo/basics/) | Covered |
| Chapter 3 Basic Data Types | Stage 1 | `go.basics` | [BasicGo/basics](../BasicGo/basics/) | Partially covered |
| Chapter 4 Composite Types | Stage 1 | `go.structs` | [BasicGo/slicemap](../BasicGo/slicemap/), [BasicGo/structs](../BasicGo/structs/) | Covered |
| Chapter 5 Functions | Stage 1 | `go.basics`, `go.errors` | [BasicGo/basics](../BasicGo/basics/), [BasicGo/defer](../BasicGo/defer/) | Covered |
| Chapter 6 Methods | Stage 1 | `go.structs` | [BasicGo/structs](../BasicGo/structs/) | Covered |
| Chapter 7 Interfaces | Stage 1 | `go.interfaces` | [BasicGo/interface](../BasicGo/interface/), [Utils/Interfaces](../Utils/Interfaces/) | Covered |
| Chapter 8 Goroutines and Channels | Stage 2 | `concurrency.goroutine`, `concurrency.channel` | [BasicGo/GoRoutine](../BasicGo/GoRoutine/), [BasicGo/channelselect](../BasicGo/channelselect/) | Covered |
| Chapter 9 Shared Variables | Stage 2 | `concurrency.sharedvars` | [BasicGo/sharedvars](../BasicGo/sharedvars/) | Covered |
| Chapter 10 Packages and the Go Tool | Stage 4 | `engineering.tooling`, `engineering.api` | [go.mod](../go.mod), [CONTRIBUTING.md](../CONTRIBUTING.md), [project standards](project-standards.md) | Documented |
| Chapter 11 Testing | Stage 1 to Stage 4 | `go.testing`, `engineering.tooling` | [BasicGo/testingdemo](../BasicGo/testingdemo/) and package tests | Covered |
| Chapter 12 Reflection | Stage 2 | `go.reflection` | [BasicGo/reflect](../BasicGo/reflect/) | Covered |
| Chapter 13 Low-Level Programming | Stage 2 to Stage 4 | `go.lowlevel`, `engineering.performance` | [BasicGo/lowlevel](../BasicGo/lowlevel/) | Introductory coverage |

## Remaining Gaps

- Add fuller examples for arrays, strings, Unicode, and byte slices.
- Add a tooling walkthrough for `go list`, `go test -run`, `go test -bench`, and `go test -race`.
- Add benchmarks for `Sorts`, `Heap`, `Union`, and `Graph_algo`.
- Expand `BasicGo/lowlevel` with clearer `unsafe` risk notes and alternatives.
