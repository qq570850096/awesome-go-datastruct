# Awesome Go Data Structures

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> 中文主入口：从 Go 初学者逐步成长为 Go 工程专家。  
> English summary: a hands-on Go learning repository for language fundamentals, data structures, algorithms, concurrency, testing, and engineering practice.

这个仓库不是单纯的代码片段合集，而是一条可执行的 Go 学习路线：先理解语言机制，再掌握数据结构与算法，最后进入并发、测试、工程化和性能意识。每个阶段都尽量配套源码、测试命令、练习任务和后续进阶方向。

## 项目定位

- 面向中文读者，提供从 Go 基础到工程专家的学习入口。
- 以数据结构和算法为核心训练场，补足 Go 语言机制、并发、测试和工程设计。
- 借鉴 TheAlgorithms/Python 的 README 门户、目录索引和贡献规范思路，但保持本仓库的 Go 工程学习定位。
- 所有学习内容都尽量落到可运行的包、测试和练习任务上。

## 适合人群

| 角色 | 建议入口 |
| --- | --- |
| Go 零基础学习者 | 从 `BasicGo/` 和 [从初学者到工程专家](docs/roadmaps/beginner-to-expert.md) 开始 |
| 有其他语言基础 | 先读 [知识谱图](docs/knowledge-graph.md)，跳过已经熟悉的通用概念 |
| 算法刷题者 | 从 `Linked/`、`Sorts/`、`Graph_algo/` 和 [目录索引](DIRECTORY.md) 开始 |
| 后端工程师 | 重点学习并发、`context`、测试、设计模式、`webdemo/` 和 [专家路线](docs/expert-track.md) |

## 30 秒快速开始

环境要求：Go 1.22 或更新版本。

```bash
go vet ./...
go test ./...
```

只运行一个主题：

```bash
go test ./BasicGo/pointers
go test ./DoubleLinked -run TestFIFO
go test ./Graph_algo/search
```

## 四阶段学习总览

| 阶段 | 目标 | 推荐模块 | 完成标志 |
| --- | --- | --- | --- |
| Stage 1 入门 | 掌握 Go 基础语法、函数、指针、结构体、接口、错误处理和测试 | `BasicGo/` | 能读懂并修改基础示例，能写表驱动测试 |
| Stage 2 进阶 | 理解 goroutine、channel、select、context、共享变量和 race 思维 | `BasicGo/GoRoutine`、`BasicGo/channelselect`、`BasicGo/context`、`BasicGo/sharedvars` | 能解释并发程序的数据流和同步边界 |
| Stage 3 核心算法 | 掌握链表、栈队列、堆、树、图、排序、并查集、Trie、SkipList | `Linked/`、`Heap/`、`Graph_algo/`、`Sorts/` | 能用测试验证结构不变量和复杂度权衡 |
| Stage 4 工程专家 | 建立测试、benchmark、vet、race、模块化、API 设计和 Web 工程意识 | `DesignPatterns/`、`webdemo/`、`OSExam/` | 能按工程标准维护和扩展模块 |

完整路线见 [docs/roadmaps/beginner-to-expert.md](docs/roadmaps/beginner-to-expert.md)。

## 学习入口

| 入口 | 用途 |
| --- | --- |
| [知识谱图](docs/knowledge-graph.md) | 看清 Go 学习概念之间的前置、练习和深化关系 |
| [结构化知识数据](docs/data/knowledge-graph.json) | 后续可用于生成网页、目录或可视化 |
| [DIRECTORY.md](DIRECTORY.md) | TheAlgorithms 风格的全仓库主题目录 |
| [学习型目录](docs/catalog.md) | 查看难度、前置知识、下一步主题和练习任务 |
| [Go 语言圣经对照](docs/go-bible-gap-map.md) | 对照 The Go Programming Language 章节覆盖情况 |
| [练习矩阵](docs/exercise-matrix.md) | 按阶段完成阅读、改写、补测试、实现、benchmark 和重构任务 |
| [专家路线](docs/expert-track.md) | 从会写 Go 走向能做 Go 工程 |
| [贡献指南](CONTRIBUTING.md) | 新增主题、测试和文档时遵循的规则 |

## 仓库模块地图

| 路径 | 学习内容 |
| --- | --- |
| `BasicGo/` | Go 语言基础、指针、并发、context、反射、泛型、unsafe 入门 |
| `Linked/`、`DoubleLinked/` | 单链表、双链表、LRU/LFU/FIFO 缓存 |
| `stack/`、`queue/`、`main/622.go` | 栈、队列、循环队列 |
| `BinarySearch/`、`AVL/`、`Red-Black/`、`Segment/` | 搜索树、平衡树、线段树 |
| `Trie/`、`skiplists/`、`Union/`、`Heap/`、`Set/` | 前缀树、跳表、并查集、堆、集合 |
| `Sorts/` | 冒泡、插入、选择、归并、快速、希尔、桶和计数排序 |
| `Graph_algo/` | 图表示、BFS、DFS、路径、环检测、二分图和图题练习 |
| `DesignPatterns/` | 创建型、结构型、行为型和复合设计模式 |
| `OSExam/` | FCFS、SJF、优先级调度和文件系统骨架 |
| `webdemo/` | HTTP、Gin、mini-Gin 和红包 demo |

## 推荐学习循环

1. 先在 [知识谱图](docs/knowledge-graph.md) 找到当前位置。
2. 按 [从初学者到工程专家](docs/roadmaps/beginner-to-expert.md) 选择阶段。
3. 进入对应目录阅读 README 和源码。
4. 运行该模块测试，例如 `go test ./Graph_algo/search`。
5. 完成 [练习矩阵](docs/exercise-matrix.md) 中的阅读、改写、补测试或实现任务。
6. 最后运行 `go vet ./...` 和 `go test ./...`。

## 质量门槛

所有改动至少保持：

```bash
go vet ./...
go test ./...
```

新增主题需要补齐：

- 主题 README 或文档说明。
- 表驱动测试。
- 复杂度与核心不变量。
- `DIRECTORY.md`、`docs/catalog.md` 和 `docs/data/knowledge-graph.json` 中的索引。

## 后续路线

下一步计划见 [ROADMAP.md](ROADMAP.md)。重点包括：自动生成目录、补充 benchmark、加入 `go test -race` 示例、处理 Dependabot 漏洞、增强工程专家路线。

## License

MIT. See [LICENSE](LICENSE).
