# Go Bible Gap Map

This map compares the repository with the chapter structure of *The Go Programming Language*.

Reference:

- Chinese mirror: <https://gopl-zh.github.io/>
- Original book site: <https://www.gopl.io/>

| Book area | Repository coverage | Status |
| --- | --- | --- |
| Tutorial | `BasicGo/basics`, package tests | Covered |
| Program structure | `BasicGo/basics`, `BasicGo/pointers` | Covered |
| Basic data types | `BasicGo/basics` | Partially covered |
| Composite types | `BasicGo/slicemap`, `BasicGo/structs` | Covered |
| Functions | `BasicGo/basics` | Covered |
| Methods | `BasicGo/structs` | Covered |
| Interfaces | `BasicGo/interface`, `Utils/Interfaces` | Covered |
| Goroutines and channels | `BasicGo/GoRoutine`, `BasicGo/channelselect`, `BasicGo/context` | Covered |
| Shared variables | `BasicGo/sharedvars` | Covered |
| Packages and go tool | `go.mod`, `CONTRIBUTING.md`, local command docs | Documented |
| Testing | `BasicGo/testingdemo`, repository-wide table tests | Covered |
| Reflection | `BasicGo/reflect` | Covered |
| Low-level programming | `BasicGo/lowlevel` | Introductory coverage |

## Remaining Opportunities

- Add examples for arrays and strings as first-class topics.
- Add a package/tooling walkthrough for `go list`, `go test`, `go vet`, and `go test -bench`.
- Add race detector notes using `go test -race`.
- Add more examples for `unsafe` tradeoffs, but keep them explicitly marked as low-level and exceptional.
