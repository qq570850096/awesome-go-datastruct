## 广度优先搜索（BFS）

### 概述

广度优先搜索从起始顶点开始，按层次逐层访问所有可达顶点。本模块实现了基于队列的 BFS 遍历。

### 特性

| 属性 | 值 |
|-----|-----|
| 时间复杂度 | O(V + E) |
| 空间复杂度 | O(V) |
| 数据结构 | 队列 |
| 遍历顺序 | 按距离从近到远 |

### 核心实现

```go
// 广度优先遍历取所有节点
func Traverse(hash *Adj.Hash, start int) (order []int) {
    visited := make([]bool, hash.V())

    bfs := func(source int) {
        que := []int{source}      // 队列初始化
        visited[source] = true

        for len(que) > 0 {
            temp := que[0]        // 取队首
            que = que[1:]         // 出队
            order = append(order, temp)

            // 遍历所有邻接点
            for _, v := range hash.LinkedVertex(temp) {
                if !visited[v] {
                    visited[v] = true
                    que = append(que, v)  // 入队
                }
            }
        }
    }

    // 从指定起点开始
    bfs(start)

    // 处理非连通图的其他连通分量
    for i := 0; i < hash.V(); i++ {
        if !visited[i] {
            bfs(i)
        }
    }
    return
}
```

### 使用示例

```go
graph := &Adj.Hash{}
graph.ReadFromFile("g.txt")

// 从顶点0开始BFS遍历
order := BFS.Traverse(graph, 0)
fmt.Println("BFS遍历顺序:", order)
```

### BFS vs DFS

| 特性 | BFS | DFS |
|-----|-----|-----|
| 数据结构 | 队列 | 栈/递归 |
| 遍历顺序 | 按层次 | 按深度 |
| 最短路径 | 天然保证（无权图） | 不保证 |
| 空间占用 | 可能较大 | 较小 |

### 应用场景

- 无权图最短路径
- 层序遍历
- 社交网络关系度数计算
- 网格迷宫最短路径

### 运行测试

```bash
go test ./Graph_algo/BFS
```
