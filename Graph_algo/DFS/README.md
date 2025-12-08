## 深度优先搜索（DFS）

### 概述

深度优先搜索从起始顶点开始，沿着一条路径尽可能深入，直到无法继续再回溯。本模块实现了基于 DFS 的连通分量检测。

### 特性

| 属性 | 值 |
|-----|-----|
| 时间复杂度 | O(V + E) |
| 空间复杂度 | O(V) |
| 数据结构 | 递归栈 |
| 遍历顺序 | 按深度优先 |

### 连通分量（CC）

连通分量是无向图中的最大连通子图。CC 结构用于检测图中有多少个独立的连通区域。

```go
// 深度优先遍历找到所有连通分量
type CC struct {
    graph   *Adj.Hash
    visited []int    // visited[v] = 分量编号，-1表示未访问
    cccount int      // 连通分量数量
}

func (C *CC) Init(graph *Adj.Hash) {
    C.graph = graph
    C.visited = make([]int, C.graph.V())

    // 初始化为未访问状态
    for i := range C.visited {
        C.visited[i] = -1
    }

    // 遍历所有顶点
    for i := 0; i < C.graph.V(); i++ {
        if C.visited[i] == -1 {
            C.Dfs(i, C.cccount)  // 发现新的连通分量
            C.cccount++
        }
    }
}

// DFS递归遍历
func (C *CC) Dfs(v int, cccountid int) {
    C.visited[v] = cccountid  // 标记所属分量

    for _, w := range C.graph.LinkedVertex(v) {
        if C.visited[w] == -1 {
            C.Dfs(w, cccountid)
        }
    }
}
```

### 使用示例

```go
graph := &Adj.Hash{}
graph.ReadFromFile("g.txt")

cc := new(DFS.CC)
cc.Init(graph)

fmt.Println("连通分量数:", cc.Cccount())
```

### DFS 应用

| 应用 | 说明 |
|-----|------|
| 连通分量检测 | 统计独立连通区域 |
| 环检测 | 判断图中是否存在环 |
| 路径查找 | 寻找两点间路径 |
| 二分图判断 | 检测图是否可二染色 |
| 拓扑排序 | 有向无环图排序 |

### 运行测试

```bash
go test ./Graph_algo/DFS
```
