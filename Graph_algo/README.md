## 图论算法（Graph Algorithms）

### 概述

本模块通过 Go 语言实现图论常用算法，涵盖图的表示方法、遍历算法和图搜索问题。

### 目录结构

| 子目录 | 功能 | 说明 |
|-------|------|------|
| [Adj](./Adj/) | 图的表示 | 邻接矩阵、邻接表、哈希表实现 |
| [BFS](./BFS/) | 广度优先搜索 | 层序遍历、最短路径 |
| [DFS](./DFS/) | 深度优先搜索 | 连通分量统计 |
| [search](./search/) | 图搜索问题 | 单源路径、环检测、二分图判断 |
| [leetcode](./leetcode/) | LeetCode 题解 | 图相关题目解答 |

### 测试数据文件

| 文件 | 说明 |
|------|------|
| g.txt | 普通无向图 |
| g2.txt | 包含环的图 |
| g2_noCycle.txt | g2 删除环后的版本 |
| notBip.txt | 非二分图示例 |

### 图的表示方法

#### 邻接矩阵（Adjacency Matrix）

使用二维数组表示顶点间的连接关系：

```go
// Matrix 邻接矩阵表示
type Matrix struct {
    v   int     // 顶点数
    e   int     // 边数
    adj [][]int // 邻接矩阵
}
```

**特点：**
- 查询边 O(1)
- 空间 O(V²)
- 适合稠密图

#### 邻接表（Adjacency List）

使用数组+链表表示：

```go
// Table 邻接表表示
type Table struct {
    v   int
    e   int
    adj [][]int // 每个顶点对应一个邻居列表
}
```

**特点：**
- 查询边 O(degree)
- 空间 O(V+E)
- 适合稀疏图

#### 哈希表实现（推荐）

本项目主要使用的实现方式：

```go
// Hash 哈希表实现的无向图
type Hash struct {
    v   int            // 顶点数
    e   int            // 边数
    adj map[int][]int  // 哈希表存储邻接关系
}

// 从文件读取图
func (hash *Hash) ReadFromFile(filename string) error {
    // 第一行：顶点数 边数
    // 后续行：边的两个顶点
    // 自动检测平行边和自环
}

// 获取顶点的所有邻居
func (hash *Hash) LinkedVertex(v int) []int

// 检测边是否存在
func (hash *Hash) HasEdge(v, e int) bool

// 计算顶点的度
func (hash *Hash) Degree(v int) int
```

### 广度优先搜索（BFS）

从起点开始，逐层向外扩展：

```go
// Traverse 广度优先遍历
func Traverse(graph *Adj.Hash, start int) []int {
    visited := make([]bool, graph.V())
    result := make([]int, 0)
    queue := []int{start}
    visited[start] = true

    for len(queue) > 0 {
        v := queue[0]
        queue = queue[1:]
        result = append(result, v)

        for _, w := range graph.LinkedVertex(v) {
            if !visited[w] {
                visited[w] = true
                queue = append(queue, w)
            }
        }
    }
    return result
}
```

**应用场景：**
- 无权图最短路径
- 层序遍历
- 社交网络中的度数分离

### 深度优先搜索（DFS）

沿着一条路径深入，回溯后探索其他路径：

```go
// DFS 深度优先遍历
func DFS(graph *Adj.Hash, v int, visited []bool, result *[]int) {
    visited[v] = true
    *result = append(*result, v)

    for _, w := range graph.LinkedVertex(v) {
        if !visited[w] {
            DFS(graph, w, visited, result)
        }
    }
}

// CC 连通分量统计
type CC struct {
    graph   *Adj.Hash
    visited []int // -1 表示未访问，否则表示所属连通分量编号
    count   int   // 连通分量数量
}

func (cc *CC) IsConnected(v, w int) bool {
    return cc.visited[v] == cc.visited[w]
}
```

**应用场景：**
- 连通性检测
- 路径查找
- 拓扑排序
- 环检测

### 图搜索问题

#### 单源路径（Single Source Path）

从源点到所有其他顶点的路径：

```go
type SingleSource struct {
    graph  *Adj.Hash
    source int       // 源点
    pre    []int     // 前驱数组，用于重建路径
}

// Path 获取从源点到目标的路径
func (ss *SingleSource) Path(target int) []int {
    if !ss.IsConnectedTo(target) {
        return nil
    }
    path := []int{}
    for cur := target; cur != ss.source; cur = ss.pre[cur] {
        path = append([]int{cur}, path...)
    }
    return append([]int{ss.source}, path...)
}
```

#### 环检测（Cycle Detection）

```go
type Cycle struct {
    graph    *Adj.Hash
    visited  []bool
    hasCycle bool
}

func (c *Cycle) Init(graph *Adj.Hash) {
    c.graph = graph
    c.visited = make([]bool, graph.V())
    c.hasCycle = false

    for v := 0; v < graph.V(); v++ {
        if !c.visited[v] {
            if c.dfs(v, v) {
                c.hasCycle = true
                return
            }
        }
    }
}

func (c *Cycle) dfs(v, parent int) bool {
    c.visited[v] = true
    for _, w := range c.graph.LinkedVertex(v) {
        if !c.visited[w] {
            if c.dfs(w, v) {
                return true
            }
        } else if w != parent {
            // 访问到已访问的非父节点，存在环
            return true
        }
    }
    return false
}
```

#### 二分图检测（Bipartite Graph）

```go
type BipartitionDetection struct {
    graph    *Adj.Hash
    visited  []bool
    colors   []int  // 0 或 1
    isBipart bool
}

// 使用染色法：相邻顶点颜色不同则为二分图
func (b *BipartitionDetection) dfs(v, color int) bool {
    b.visited[v] = true
    b.colors[v] = color

    for _, w := range b.graph.LinkedVertex(v) {
        if !b.visited[w] {
            if !b.dfs(w, 1-color) {
                return false
            }
        } else if b.colors[w] == color {
            return false
        }
    }
    return true
}
```

### 图文件格式

```
7 6          // 第一行：顶点数 边数
0 1          // 后续行：每条边的两个顶点
0 2
1 3
1 4
2 3
2 6
```

### 测试用例

```go
func TestGraph(t *testing.T) {
    // 创建图
    graph := &Adj.Hash{}
    graph.ReadFromFile("g.txt")

    // 打印图信息
    fmt.Println(graph) // V = 7, E = 6 ...

    // BFS 遍历
    result := BFS.Traverse(graph, 0)
    fmt.Println("BFS:", result) // [0 1 2 3 4 6]

    // 环检测
    cycleDetector := &search.Cycle{}
    cycleDetector.Init(graph)
    fmt.Println("有环:", cycleDetector.HasCycle())

    // 二分图检测
    bipart := &search.BipartitionDetection{}
    bipart.Init(graph)
    fmt.Println("是二分图:", bipart.IsBipart())
}
```

### 运行方式

```bash
go run ./Graph_algo/main.go
```

### 复杂度总结

| 算法 | 时间复杂度 | 空间复杂度 | 说明 |
|------|-----------|-----------|------|
| BFS | O(V+E) | O(V) | 广度优先 |
| DFS | O(V+E) | O(V) | 深度优先 |
| 连通分量 | O(V+E) | O(V) | 基于 DFS |
| 环检测 | O(V+E) | O(V) | 基于 DFS |
| 二分图 | O(V+E) | O(V) | 染色法 |

### LeetCode 实战

#### [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

使用 DFS/BFS 统计连通分量：

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    m, n := len(grid), len(grid[0])
    count := 0

    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
            return
        }
        grid[i][j] = '0' // 标记已访问
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                count++
                dfs(i, j)
            }
        }
    }
    return count
}
```

#### [785. 判断二分图](https://leetcode-cn.com/problems/is-graph-bipartite/)

染色法判断：

```go
func isBipartite(graph [][]int) bool {
    n := len(graph)
    colors := make([]int, n)
    for i := range colors {
        colors[i] = -1
    }

    var dfs func(node, color int) bool
    dfs = func(node, color int) bool {
        colors[node] = color
        for _, neighbor := range graph[node] {
            if colors[neighbor] == -1 {
                if !dfs(neighbor, 1-color) {
                    return false
                }
            } else if colors[neighbor] == color {
                return false
            }
        }
        return true
    }

    for i := 0; i < n; i++ {
        if colors[i] == -1 {
            if !dfs(i, 0) {
                return false
            }
        }
    }
    return true
}
```

#### [207. 课程表](https://leetcode-cn.com/problems/course-schedule/)

有向图环检测（拓扑排序）：

```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    // 构建邻接表
    graph := make([][]int, numCourses)
    inDegree := make([]int, numCourses)

    for _, pre := range prerequisites {
        graph[pre[1]] = append(graph[pre[1]], pre[0])
        inDegree[pre[0]]++
    }

    // BFS 拓扑排序
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    count := 0
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        count++

        for _, neighbor := range graph[node] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    return count == numCourses
}
```
