# Directory

This is the manually maintained, TheAlgorithms-style index for the repository. It is grouped by learning topic rather than by filesystem order.

## Core Textbook Chapters

These chapters form the first complete learning spine. Each one contains exercises with hints and reference answers.

| Area | Chapters |
| --- | --- |
| Go fundamentals | [BasicGo](BasicGo/) |
| Linear structures | [Linked](Linked/), [DoubleLinked](DoubleLinked/), [stack](stack/), [queue](queue/), [Circular queue](main/) |
| Priority and connectivity | [Heap](Heap/), [Union](Union/) |
| Trees and search structures | [BinarySearch](BinarySearch/), [AVL](AVL/), [Red-Black](Red-Black/), [Segment](Segment/), [Trie](Trie/), [skiplists](skiplists/) |
| Sorting | [Sorts](Sorts/) |
| Graph algorithms | [Graph_algo](Graph_algo/), [Adj](Graph_algo/Adj/), [BFS](Graph_algo/BFS/), [DFS](Graph_algo/DFS/), [search](Graph_algo/search/), [leetcode](Graph_algo/leetcode/) |

## Go Basics

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [BasicGo/basics](BasicGo/basics/) | Beginner | Syntax, control flow, functions, constants | None | `go test ./BasicGo/basics` |
| [BasicGo/pointers](BasicGo/pointers/) | Beginner | Address semantics, mutation through pointers | Basic syntax | `go test ./BasicGo/pointers` |
| [BasicGo/slicemap](BasicGo/slicemap/) | Beginner | Slices, maps, shared backing arrays, sets | Basic syntax | `go test ./BasicGo/slicemap` |
| [BasicGo/structs](BasicGo/structs/) | Beginner | Structs, methods, embedding, tags | Pointers | `go test ./BasicGo/structs` |
| [BasicGo/interface](BasicGo/interface/) | Beginner | Interfaces, `any`, custom errors | Structs and methods | `go test ./BasicGo/interface` |
| [BasicGo/errors](BasicGo/errors/) | Beginner | Error wrapping, `errors.Is`, `errors.As` | Interfaces | `go test ./BasicGo/errors` |
| [BasicGo/testingdemo](BasicGo/testingdemo/) | Beginner | Table-driven tests, benchmarks | Functions | `go test ./BasicGo/testingdemo` |

## Concurrency

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [BasicGo/GoRoutine](BasicGo/GoRoutine/) | Intermediate | Goroutines, mutexes, wait groups, object pools | Basic Go | `go test ./BasicGo/GoRoutine` |
| [BasicGo/channelselect](BasicGo/channelselect/) | Intermediate | Channels, fan-in, tickers, `select` | Goroutines | `go test ./BasicGo/channelselect` |
| [BasicGo/context](BasicGo/context/) | Intermediate | Timeout, cancellation, request IDs | Channels | `go test ./BasicGo/context` |
| [BasicGo/sharedvars](BasicGo/sharedvars/) | Intermediate | Mutex, atomic, once, race thinking | Goroutines | `go test ./BasicGo/sharedvars` |

## Linear Structures

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [Linked](Linked/) | Advanced | Singly linked list, reversal, duplicate removal | Pointers | `go test ./Linked` |
| [DoubleLinked](DoubleLinked/) | Advanced | Doubly linked list, LRU, LFU, FIFO | Linked lists and maps | `go test ./DoubleLinked` |
| [stack](stack/) | Beginner | LIFO ordering | Slices | `go test ./stack` |
| [queue](queue/) | Beginner | FIFO ordering | Slices | `go test ./queue` |
| [main/622.go](main/622.go) | Beginner | Circular queue | Arrays and indexes | `go test ./main` |

## Trees And Search Structures

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [BinarySearch](BinarySearch/) | Advanced | BST ordering, traversal, deletion | Recursion | `go test ./BinarySearch` |
| [AVL](AVL/) | Advanced | Height balance and rotations | BST | `go test ./AVL` |
| [Red-Black](Red-Black/) | Advanced | Color invariants and rotations | BST | `go test ./Red-Black` |
| [Segment](Segment/) | Advanced | Range aggregation and update | Recursion | `go test ./Segment` |
| [Trie](Trie/) | Advanced | Prefix sharing | Strings and maps | `go test ./Trie` |
| [skiplists](skiplists/) | Advanced | Randomized layered lists | Linked lists | `go test ./skiplists` |
| [Set](Set/) | Intermediate | Set behavior with list and BST forms | Lists and trees | `go test ./Set` |

## Graph Algorithms

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [Graph_algo/Adj](Graph_algo/Adj/) | Advanced | Matrix, list, and hash graph representations | Slices and maps | `go test ./Graph_algo/Adj` |
| [Graph_algo/BFS](Graph_algo/BFS/) | Advanced | Breadth-first traversal | Queue | `go test ./Graph_algo/BFS` |
| [Graph_algo/DFS](Graph_algo/DFS/) | Advanced | Depth-first traversal and connected components | Recursion | `go test ./Graph_algo/DFS` |
| [Graph_algo/search](Graph_algo/search/) | Advanced | Paths, cycle detection, bipartite checks | BFS and DFS | `go test ./Graph_algo/search` |
| [Graph_algo/leetcode](Graph_algo/leetcode/) | Advanced | Applied graph search problems | Graph traversal | `go test ./Graph_algo/leetcode` |

## Sorting

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [Sorts](Sorts/) | Advanced | Bubble, insertion, selection, merge, quick, shell, bucket, counting | Arrays and recursion | `go test ./Sorts` |

## Design Patterns

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [DesignPatterns/CreativeType](DesignPatterns/CreativeType/) | Expert | Singleton, prototype, factories, builder | Interfaces | `go test ./DesignPatterns/CreativeType` |
| [DesignPatterns/StructuralType](DesignPatterns/StructuralType/) | Expert | Adapter, proxy, decorator, facade, bridge, composite, filter, flyweight | Interfaces and composition | `go test ./DesignPatterns/StructuralType` |
| [DesignPatterns/BehavioralType](DesignPatterns/BehavioralType/) | Expert | Mediator, command, memento, template, state, strategy, observer, interpreter, chain, iterator | Interfaces and responsibility boundaries | `go test ./DesignPatterns/BehavioralType` |
| [DesignPatterns/Compound.go](DesignPatterns/Compound.go) | Expert | Pattern composition in a simulation | Multiple pattern families | `go test ./DesignPatterns` |

## Systems And Web Demos

| Module | Level | Core concepts | Prerequisites | Test command |
| --- | --- | --- | --- | --- |
| [OSExam](OSExam/) | Expert | Scheduling and file-system skeletons | Sorting and state modeling | `go test ./OSExam` |
| [webdemo/http_basic](webdemo/http_basic/) | Intermediate | Native `net/http`, middleware, REST-style handlers | Basic Go and errors | `go test ./webdemo/http_basic` |
| [webdemo/minigin](webdemo/minigin/) | Expert | Router tree, middleware chain, JSON binding | HTTP and trees | `go test ./webdemo/minigin` |
| [webdemo/gin_example](webdemo/gin_example/) | Intermediate | Gin demo structure | HTTP basics | Manual demo |
