# Basic Go Practice

This module collects small runnable examples for Go language mechanics. Each subdirectory focuses on one concept and includes tests so learners can run the topic directly.

| Path | Topic | Command |
| --- | --- | --- |
| [basics](basics/) | Variables, constants, control flow, and functions | `go test ./BasicGo/basics` |
| [pointers](pointers/) | Address semantics and pointer mutation | `go test ./BasicGo/pointers` |
| [structs](structs/) | Structs, methods, embedding, and tags | `go test ./BasicGo/structs` |
| [interface](interface/) | Interfaces, dynamic dispatch, and custom errors | `go test ./BasicGo/interface` |
| [slicemap](slicemap/) | Slices, maps, backing arrays, and sets | `go test ./BasicGo/slicemap` |
| [defer](defer/) | Deferred cleanup, panic, and recover | `go test ./BasicGo/defer` |
| [GoRoutine](GoRoutine/) | Goroutines, mutexes, wait groups, and object pools | `go test ./BasicGo/GoRoutine` |
| [channelselect](channelselect/) | Channels, fan-in, tickers, and select | `go test ./BasicGo/channelselect` |
| [context](context/) | Timeout, cancellation, and request IDs | `go test ./BasicGo/context` |
| [sharedvars](sharedvars/) | Mutexes, atomics, once, and data-race thinking | `go test ./BasicGo/sharedvars` |
| [errors](errors/) | Error wrapping and inspection | `go test ./BasicGo/errors` |
| [generics](generics/) | Generic containers and functions | `go test ./BasicGo/generics` |
| [reflect](reflect/) | Runtime type inspection | `go test ./BasicGo/reflect` |
| [lowlevel](lowlevel/) | Size, alignment, offsets, and unsafe boundaries | `go test ./BasicGo/lowlevel` |
| [testingdemo](testingdemo/) | Table-driven tests and benchmarks | `go test ./BasicGo/testingdemo` |

## Learning Order

1. Start with `basics`, `pointers`, and `testingdemo`.
2. Move to `structs`, `interface`, `slicemap`, and `errors`.
3. Study `GoRoutine`, `channelselect`, `context`, and `sharedvars`.
4. Use `generics`, `reflect`, and `lowlevel` as advanced language tools.

## Quality Gate

```bash
go test ./BasicGo/...
```
