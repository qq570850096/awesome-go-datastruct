# Go 数据结构与工程实践学习总结

## 总览
本仓库围绕 “数据结构 + Go 特性 + 工程化能力” 展开，除常见链表、树、图、并查集、堆等结构外，还包含并发、调度算法、设计模式及实用案例。学习路径可分为三层：语言与并发基础、核心数据结构与算法、工程与架构专题。以下按模块给出要点与实践经验。

---

## 1. 基础能力模块

### 1.1 Go 语言特性 (`BasicGo`)
| 示例/算法 | 目标 | 关键点 | 文件 |
| --- | --- | --- | --- |
| 反射类型检查 | 运行时判别类型 | `reflect.TypeOf` + `Kind()` + `switch`，输出不同描述 | `reflect/Reflect.go` |
| `defer` + `recover` | Panic 捕获 | `defer` 内调用 `recover`，打印错误后继续执行，演示异常防护 | `defer/defer.go` |
| `Thread` vs `ThreadWrong` | 正确启动 goroutine | 对比闭包捕获循环变量的坑，正确写法把变量作为入参传入匿名函数 | `BasicGo/GoRoutine/Goroutine.go` |
| `CounterWrong` vs `Counter` | 互斥锁 | 无锁计数导致竞态；`sync.Mutex` 包裹临界区 | 同上 |
| `WaitGroupExam` | 协同行为 | 展示 `WaitGroup` + `Mutex` 统计并发计数，替代 `time.Sleep` 粗暴等待 | 同上 |
| `AsnyService` | 异步调用 | 带缓冲 channel 返回结果，`select` + `time.After` 控制超时 | 同上 |
| `Producer/Consumer` | CSP 并发模型 | 多消费者从 channel 中读取，`Producer` 负责关闭通道 | 同上 |
| `Cancel` | Context 取消 | `context.WithCancel` 派发取消信号，各 goroutine 调用 `isCancelledWithctx` 退出 | 同上 |
| `ObjPool` | 对象池 | 缓冲 channel 预加载对象，`GetObj` 支持超时，`ReleaseObj` 防止溢出 | 同上 |

### 1.2 通用接口 (`Utils/Interfaces`)
- 利用 `reflect.TypeOf` 确认类型，再通过 `type switch` 执行比较。
- 实践经验：若要支持更多类型，可通过泛型或接口约束，避免频繁 `panic`。

---

## 2. 线性结构与缓存模块

### 2.1 单链表 (`Linked`)
| 算法/练习 | 目标 | 关键步骤 | 文件 |
| --- | --- | --- | --- |
| `AddIndex` / `Remove` | 任意位置插入/删除 | 通过 `dummyHead` 统一处理首节点，遍历 index 次定位前驱，再插入/断链并维护 `size` | `List.go:43-125` |
| `Sort` (链表快排) | 原地排序 | `qsortList` 采用 `[low, high)` 范围；`partitionList` 用首节点为 pivot，遍历时通过交换节点值实现 | `List.go:149-177` |
| `1.1 Reverse` | 迭代反转 | 三指针 (`pre/cur/next`) 逐节点反转；`RecursiveReverse` 演示递归回溯；`InsertReverse` 采用头插法 | `1.1 Reverse.go` |
| `1.2 RemoveDup` | 删除重复节点 | 使用哈希或双指针；重点在维护前驱并跳过重复节点 | `1.2 RemoveDup.go` |
| `1.3 AddTwoList` | 两数相加 | 同步遍历两链表，维护进位，新建结果链 | `1.3 AddTwoList.go` |
| `1.4 MiddleReverse` | 部分反转 | 快慢指针定位中点，再局部反转 | `1.4MiddleReverse.go` |
| `1.5 FindLastK` | 倒数第 k 节点 | 双指针，second 先走 k，再同步走 | `1.5 FindLastK.go` |
| `1.6 FindEntry` | 检测环入口 | 快慢指针相遇后，一个回到头，同速前进即可找到入口 | `1.6 FindEntry.go` |
| `1.7 ReverseTwo` | 成对反转 | 每次处理两个节点，使用 dummyHead 便于连接 | `1.7 ReverseTwo.go` |
| `1.8 ReverseK` | 前 k 个反转 | 先截断第 k+1 个节点，再复用反转逻辑 | `1.8 ReverseK.go` |
| `1.9 MergeTwoList` | 合并有序链表 | 双指针比较后拼接 | `1.9 MeegeTwoList.go` |
| `1.10 RemoveNode` | O(1) 删除节点 | 将待删节点值替换为 next，再跳过 next（尾节点需要顺序遍历） | `1.10 RemoveNode.go` |
| `1.11 CheckIntersect` | 判断链表是否相交 | 分别计算长度并对齐，再同步前进比对节点地址 | `1.11 CheckIntersect.go` |

### 2.2 双向链表与缓存 (`DoubleLinked`)
| 算法/策略 | 关键思想 | 实现细节 | 文件 |
| --- | --- | --- | --- |
| LRU | 最近最少使用淘汰 | 哈希存键→节点，双向链表记录访问顺序：命中移动到头，满时淘汰尾节点 | `LRU.go` |
| LFU | 最不常使用淘汰 | `find` 维护 `key→LFUNode(freq + Node)`；`freq_map` 映射频率到链表，淘汰最小频率链表的尾节点；更新频率时需移动节点 | `LFU.go` |
| FIFO | 先进先出 | 仅用链表的头尾表示先入先出顺序，缺页直接替换 | `FIFO.go` |
| 双向链表基类 | 节点增删 | `List.go` 提供 `Append/AppendToHead/Remove/Pop` 等操作，所有缓存策略共享 | `List.go` |

要点：缓存实现强调 **map + 双向链表** 协同；LFU 需额外频率桶；所有写操作需保持 `size` 一致，避免内存泄漏。

### 2.3 栈与队列
| 算法 | 功能 | 核心步骤 | 文件 |
| --- | --- | --- | --- |
| 顺序栈 + 括号匹配 | 判断括号有效性 | Push 左括号，遇右括号弹栈并检查匹配；遍历结束栈空即合法 | `stack/stack.go` |
| 普通循环队列 | 顺序队列，示例 `main` | `front/rear` 用模运算，`size` 判断空满；示例 `main` 中循环入队/出队 | `queue/queue.go` |
| 力扣 622 循环队列 | 接口化实现 | `EnQueue/DeQueue/Front/Rear/IsFull/IsEmpty` 按 leetcode 定义实现，并采用模运算处理索引 | `main/622.go` |

---

## 3. 树结构与高级数据结构

### 3.1 二叉搜索树 (`BinarySearch`)
| 算法 | 功能 | 关键实现 | 文件 |
| --- | --- | --- | --- |
| `AddE` / `add` | 插入节点 | 递归定位空位，`size++`；若允许重复可以调整左右子树比较条件 | `Tree.go:62-101` |
| `Contains` | 查找 | 递归比较值，直到命中或空节点 | `Tree.go:107-128` |
| 遍历（前、中、后、层序、非递归前序） | 输出结构 | 通过递归或手动栈/队列实现 | `Tree.go:134-189` |
| `FindMin` / `FindMax` | Extremum | 顺着左/右子树走到叶子 | `Tree.go:190-214` |
| `DelMin` / `DelMax` | 删除极值 | `rmMin`/`rmMax` 递归返回新根，并维护 `size` | `Tree.go:215-220`、`Tree.go:226-220`? (typo) |
| `Remove` | 删除指定值 | 分左右子树空/非空三种情况；当两侧皆存在时使用后继节点（右子树最小值）替换 | `Tree.go:221-253` |
| `String` | 打印结构 | 使用深度前缀 `--` 可视化树形结构 | `Tree.go:255-274` |

### 3.2 红黑树 (`Red-Black`)
| 功能 | 关键步骤 | 文件 |
| --- | --- | --- |
| 插入 (`Push`) | 永远插入红节点；递归插入后依次检查「右红但左黑 → 左旋」「左红且左左红 → 右旋」「左右皆红 → 颜色翻转」，最后把根设为黑 | `Red-Black/Tree.go:39-118` |
| 旋转 (`leftRotate`/`rightRotate`) | 维护 2-3 树等价结构，旋转时传递颜色 | `Tree.go:119-156` |
| 查找/更新 (`getNode`, `Contains`, `GetValue`, `SetNewValue`) | 与 BST 类似，但需保持树平衡 | `Tree.go:158-204` |
| 删除 (`Remove`) | 目前实现基于 BST 的非严格红黑删除逻辑；可在此基础上扩展左倾红黑树的删除修正 | `Tree.go:206-240` |

### 3.3 线段树 (`Segment`)
| 算法 | 功能 | 关键实现 | 文件 |
| --- | --- | --- | --- |
| `Init` | 构造线段树 | 拷贝原数组，`tree` 开 4 倍空间，调用 `buildSegmentTree` | `Tree.go:17-38` |
| `buildSegmentTree` | 递归建树 | 叶子存原值，内部节点通过 `merger(left, right)` 聚合 | `Tree.go:31-43` |
| `QueryLR` | 区间查询 | 检查越界后调用 `query`，`query` 根据区间分布决定下探左/右或者拆分求并 | `Tree.go:45-69` |
| `Update` | 单点更新 | 更新 `data[index]`，递归到叶子后回溯重新聚合 | `Tree.go:71-80` |
| 辅助接口 | `leftChild/rightChild/GetSize/Get` | 提供安全访问和调试输出 | `Tree.go:82-118` |

### 3.4 并查集 (`Union`)
| 版本 | 核心思想 | 关键函数 | 文件 |
| --- | --- | --- | --- |
| Quick-Find (`Find` struct) | 数组 `id[i]` 记录连通分量 id，`Union` 需要遍历整个数组改写 id，适合演示 | `Find.go:6-53` |
| Quick-Union (`QuickFind` struct) | `parent[i]` 指向父节点，通过连续 `parent[p]` 找根；`Union` 将一颗树根接到另一颗树，易与路径压缩结合 | `Find.go:54-110` |
| 公共接口 | `Union(p,q)`、`IsConnect(p,q)`、`GetSize()` | `Union` 逻辑实现 | `Find.go` |

### 3.5 堆与优先队列 (`Heap`)
| 算法 | 目标 | 关键实现 | 文件 |
| --- | --- | --- | --- |
| `Add` + `siftUp` | 上浮构建大顶堆 | 新元素插入末尾后与父节点比较、交换 | `Array.go:40-61` |
| `RemoveMax` + `siftDown` | 删除堆顶 | 将最后元素换到根，下沉以恢复堆序 | `Array.go:63-89` |
| `Replace` | 复用堆顶 | 直接覆盖后下沉 | `Array.go:64-69` |
| `InitHeapWithArray` | Heapify | 自底向上调用 `siftDown`，复杂度 O(n) | `Array.go:7-26` |
| `PriorityQueue` | 面向接口的优先队列 | 组合 `MaxHeap` 实现 `Enqueue/Dequeue/GetFront/Size/IsEmpty` | `PriorityQueue.go` |
| 性能测试 | 压测堆操作 | 随机数插入/删除 100 万次验证正确性与性能 | `Arrar_test.go` |

### 3.6 排序算法 (`Sorts`)
| 算法 | 核心思想 | 关键实现/注意事项 | 文件 |
| --- | --- | --- | --- |
| 冒泡排序 | 相邻比较交换，越大的元素冒泡到末尾 | 双循环，逐渐缩短未排序区间 | `Bubble.go` |
| 插入排序 | 构建有序序列，将新元素插入合适位置 | 内层 while 右移元素给插入值腾位置 | `Insert.go` |
| 选择排序 | 每轮选择最小值放到开头 | 双循环选择最小值索引，再交换 | `Select.go` |
| 希尔排序 | 分组插入排序，加快局部有序 | gap 从 `n/2` 逐步减半，子序列内做插入排序 | `Shell.go` |
| 归并排序 | 分治合并 | `Merge.go` / `MergeBottomUp.go`（若有）实现递归与迭代，重点在临时数组合并 | `Merge.go` |
| 快速排序（单路/双路/三路） | 划分 + 递归 | `Quick.go` 中 `partition`/`partition2`/`quickSort3Ways`；随机 pivot 避免退化 | `Quick.go` |
| 桶排序 | 计数/桶分布 | 根据值区间映射到桶，桶内排序后拼接 | `Bucket.go` |
| 基准测试 | 通过 `testing` 验证 | `Sort_test.go` 使用 `t.Log` 展示结果，可扩展为 `Benchmark` | `Sort_test.go` |

### 3.7 Trie 与 SkipList
| 数据结构 | 功能 | 实现细节 | 文件 |
| --- | --- | --- | --- |
| Trie | 字典树 + 模糊匹配 | `Push` 按字符逐层创建节点；`Contains` 严格匹配；`SearchPrefix` 检测前缀；`MatchSearch` 支持 `.` 任意字符匹配（递归遍历所有子节点） | `Trie.go` |
| Trie 统计 | 单词去噪统计 | `Trie_test.go` 读取文本 -> 标准化词 -> 插入 Trie，统计单词数量与耗时 | `Trie_test.go` |
| SkipList | 多层有序链 | `Insert` 使用 `randomLevel` 提升高度，`forward` 数组保存多层指针；`Search` 从最高层逐渐往下；`Delete` 维护 `update` 路径并在必要时降低 `level` | `skiplists/SkipLists.go` |

---

## 4. 图算法模块 (`Graph_algo`)

### 4.1 图表示 (`Adj`)
| 结构 | 适用场景 | 关键实现 | 文件 |
| --- | --- | --- | --- |
| 邻接矩阵 (`Matrix`) | 稠密图，快速判断边 | 使用二维切片表示 `graph[v][w]`，读取文件时检查自环/平行边 | `Adj/Matrix.go` |
| 邻接表 (`Table`) | 稀疏图或无向图 | `adj [][]int` 存储边列表；`ReadFromFile` 逐行添加双向边 | `Adj/Table.go` |
| 哈希邻接表 (`Hash`) | 灵活扩容，用 map 加速 | `adj map[int][]int` 保存邻接点；提供 `HasEdge`、`LinkedVertex`、`Degree` 等方法，`ReadFromFile` 负责格式校验 | `Adj/Hash.go` |

### 4.2 搜索与图性质
| 算法 | 目标 | 关键步骤 | 文件 |
| --- | --- | --- | --- |
| BFS 全图遍历 | 获取所有节点访问序 | 每个连通分量独立 BFS，使用队列与 `visited` 标记；可作为最短路基础 | `BFS/traverse.go` |
| DFS 连通分量 (`DFS/CC.go`) | 统计 `CC` | 从未访问节点发起 DFS，`ccId` 数组记录分量编号 | `DFS/CC.go` |
| 单源路径 (`search/SingleSourcePath.go`) | 从源点 `s` 到任意 `t` 的路径 | DFS 过程中记录 `pre` 数组，`Path(t)` 通过回溯构造路径 | `search/SingleSourcePath.go` |
| 指定终点路径 (`search/Find.go`) | 判断 `s` 到 `T` 是否可达并返回路径 | 基于 `Path` 结构，`Dfs` 提前返回；`IsConnectedTo` 查看 visited | `search/Find.go` |
| 环检测 (`search/Cycle.go`) | 判断无向图含环 | DFS 遍历，若遇到已访问且非父节点的邻居即存在环 | `search/Cycle.go` |
| 二分图检测 (`search/BipartitionDetection.go`) | 着色判断 | DFS 时给相邻节点赋予不同颜色，若出现同色邻接则非二分图 | `search/BipartitionDetection.go` |

### 4.3 LeetCode & 实战算法 (`Graph_algo/leetcode`)
| 题目 | 思路概述 | 文件 |
| --- | --- | --- |
| `4LWater`（装 4 升水问题） | BFS/DFS 状态搜索，分析水桶倒水组合 | `leetcode/4LWater.go` |
| `1091` 最短路径 | BFS 在网格上，队列存坐标与距离 | `leetcode/1091.go` |
| `695` 岛屿面积 | DFS/BFS 遍历岛屿计算面积 | `leetcode/695.go`, `695_2.go` |
| `200` 岛屿数量 | DFS 标记访问过的陆地 | `leetcode/200.go` |
| `785` 二分图 | BFS/DFS 着色，与 `BipartitionDetection` 类似 | `leetcode/785.go` |
| `752` 打开转盘锁 | BFS + visited + deadends 过滤 | `leetcode/752.go` |
| `shuffle` | Fisher-Yates 洗牌 | `leetcode/shuffle.go` |

### 4.4 其他 BFS/DFS 实战
- `Graph_algo/DFS/CC.go`、`Graph_algo/BFS/traverse.go` 可作为模板复制到新题。
- `Graph_algo/leetcode` 目录给出对应测试 (`*_test.go`)，可直接运行验证。

---

## 5. 设计模式与架构专题 (`DesignPatterns`)

### 5.1 `DesignPatterns/7Rules.go`
- 结合 SOLID、KISS、DRY、YAGNI、迪米特法则等七大设计原则，配合中文文档（`七大设计原则.md`）作为理论基础。

### 5.2 创建型模式 (`CreativeType`)
| 模式 | 关键点 | 文件 |
| --- | --- | --- |
| 简单工厂 | 静态方法根据入参返回不同产品，集中创建逻辑 | `SimpleFactory.go` |
| 抽象工厂 | 定义产品族接口，具体工厂实现多个产品的创建 | `AbstractFactory.go` |
| 单例 | 通过 `sync.Once` 或包级变量确保唯一实例 | `Singleton.go` |
| 原型 | 实现 `Clone` 方法，通过拷贝构造新对象 | `Prototype.go` |
| 建造者 | 分离复杂对象的构建与表示，Director 负责组装 | `Builder.go` |

### 5.3 结构型模式 (`StructuralType`)
| 模式 | 核心算法 | 文件 |
| --- | --- | --- |
| 适配器 | 通过包装类将不兼容接口转化 | `Wrapper.go` |
| 桥接 | 抽象与实现分离，通过组合在运行时确定 | `Bridge.go` |
| 组合 | 树状结构统一处理叶子/组合节点 | `Component.go`, `Composite` 实现 |
| 装饰器 | 动态叠加职责，保持接口一致 | `Decorator.go` |
| 外观 | 封装复杂子系统接口 | `Facade.go` |
| 享元 | 共享内部状态，减少对象数量 | `Flyweight.go` |
| 过滤器 | 组合多种过滤条件 | `Filter.go` |
| 代理 | 控制访问、增加缓存/安全层 | `Proxy.go` |

### 5.4 行为型模式 (`BehavioralType`)
| 模式 | 关键流程 | 文件 |
| --- | --- | --- |
| 命令 | 将请求封装成对象，实现可撤销队列 | `Command.go` |
| 解释器 | 定义表达式抽象，递归解释 AST | `Interpreter.go` |
| 迭代器 | 提供统一遍历接口 | `Iterator.go` |
| 观察者 | 发布-订阅，`Observer` 接收事件 | `Observer.go` |
| 状态 | 将状态封装为对象，Context 委托处理 | `State.go` |
| 策略 | 定义算法族，运行时可替换 | `Strategy.go` |
| 模板方法 | 固定流程骨架，子类填充步骤 | `Template.go` |
| 责任链 | 将请求沿链传递直至被处理 | `Handler.go` |
| 中介者 | 用中介控制对象交互，降低耦合 | `Mediator.go` |
| 备忘录 | 保存状态快照，支持撤销 | `Memento.go` |

### 5.5 复合模式示例 (`Compound.go`)
- `QuackAble` 表示鸭子接口；`GooseAdapter` 适配鹅；`QuackCounter` 装饰器记录次数；`DuckFactory` 结合抽象工厂与装饰器；`Flock` 组合多个鸭子；`ObservableAssist` + `DuckDoctor` 实现观察者。
- 该示例贯穿多个模式，展示如何将算法/模式组合满足复杂需求。

---

## 6. 系统专题与综合实践

### 6.1 调度算法 (`OSExam`)
| 算法 | 描述 | 核心流程 | 文件 |
| --- | --- | --- | --- |
| FCFS（先来先服务） | 按提交时间排序作业 | `InitFromFile` 读作业 -> `sort.Sort` -> 逐个计算开始/结束/等待时间 | `FCFS.go` |
| SJF / FP / FCFS Tests | 不同调度策略 | 通过 `Process` 结构记录 `submitTime/runTime/priority`，测试文件模拟输入 | `SJF.go`, `FP.go`, `FCFS_test.go` |
| 指标计算 | 周转/带权周转 | `Process.GetColTime/GetColTimeWithWeight` | `Process.go` |
| 文件系统骨架 | 目录/文件索引 | `UFD`/`DIR`/`Cuse` 结构展示链式目录、一级/二级索引思路 | `fileSystem.go` |

### 6.2 组合数学与其他
- `combine.go` 的 `combinationSum3`：回溯 + 剪枝 -> 使用 `temp` 存路径、`used` 标记数字、`remain` 控制剩余和；测试在 `com_test.go`。
- `main/main.go`：最小入口，可根据需要组织 CLI。

---

## 7. 问题与改进建议

1. **模块化与依赖**：仓库无 `go.mod`，导致 `go test ./...` 失败。建议在根目录初始化 `go.mod` 并调整 import 路径（如 `algo/Graph_algo/Adj`），可提升可维护性。
2. **测试体系**：部分模块已有 `*_test.go`，但许多数据结构（如 `DoubleLinked`、`Set`、`Union`）缺测试。建议新增 table-driven tests 覆盖典型场景与边界。
3. **性能与错误处理**：图算法中多处 `panic` 处理，实际工程中需返回 error；缓存模块可增加命中率统计接口、并添加并发安全策略。
4. **文档深度**：现有 README 只列结构清单，可在 `STUDY_SUMMARY.md` 基础上扩展为教程式文档，附带示意图或时序说明。
5. **示例优化**：`Trie_test` 读大文件时要考虑分词过滤与 Stop Words；`OSExam` 可引入配置文件驱动输入，增强可扩展性。

---

## 8. 学习路线建议

1. **Go 语言与并发**：先阅读 `BasicGo`，掌握 `defer/recover`、goroutine、channel、`sync`、CSP 模式。
2. **线性结构与缓存**：实现链表/栈/队列，练习常见题；再学习 LRU/LFU/FIFO 缓存如何结合哈希与双向链表。
3. **树与集合**：依次手写 BST → AVL/红黑树 → Heap → Segment Tree → Union-Find → SkipList → Trie，体会不同平衡手段。
4. **图论算法**：从图的三种表示入手，掌握 BFS/DFS、环检测、二分图、最短路（可扩展 Dijkstra/Bellman-Ford）。
5. **设计模式与工程思维**：按照创建型→结构型→行为型→复合模式顺序学习，并结合 `Compound.go` 练习模式整合。
6. **系统专题**：最后实践操作系统调度、文件系统骨架、组合算法等，提升针对具体业务的建模能力。

---

## 9. 结语

通过本项目可以系统性掌握 Go 在数据结构、算法与工程模式中的常见写法。学习过程中建议：
- 坚持 KISS/YAGNI：面向当前需求实现，逐步抽象，避免过度设计。
- 遵循 SOLID/DRY：每个结构维护单一职责，复用通用逻辑（如链表操作、频率桶管理）。
- 注重测试与文档：在完成每个模块后补充测试案例与注释，形成可持续演进的知识库。

此文档可作为进一步扩展与复习的索引，后续若新增模块（如动态规划、分布式系统等），可按相同结构补充。
