## 图的表示方法

### 概述

本目录实现了无向图的三种存储方式，提供统一的接口进行图的基本操作。

### 实现方式

| 类型 | 文件 | 空间复杂度 | 查询边 | 遍历邻接点 |
|-----|------|-----------|--------|-----------|
| 邻接矩阵 | `Matrix.go` | O(V²) | O(1) | O(V) |
| 邻接表 | `Table.go` | O(V+E) | O(degree) | O(degree) |
| 哈希表 | `Hash.go` | O(V+E) | O(1) | O(degree) |

### 统一接口

```go
// 图的基本操作
V() int                      // 返回顶点数
E() int                      // 返回边数
HasEdge(v, e int) bool       // 判断边是否存在
LinkedVertex(v int) []int    // 返回顶点v的所有邻接点
Degree(v int) int            // 返回顶点v的度数
ReadFromFile(filename string) error  // 从文件读取图数据
```

### 数据结构

#### 邻接矩阵

```go
type Matrix struct {
    v   int        // 顶点数
    e   int        // 边数
    adj [][]int    // V×V 矩阵，adj[i][j]=1 表示存在边(i,j)
}
```

#### 哈希表

```go
type Hash struct {
    v   int              // 顶点数
    e   int              // 边数
    adj map[int][]int    // 邻接表，key为顶点，value为邻接点列表
}
```

### 文件格式

图数据文件格式（如 `g.txt`）：
```
7 8        // 第一行：顶点数 边数
0 1        // 后续每行：一条边的两个端点
0 2
1 3
1 4
...
```

### 使用示例

```go
// 创建并读取图
graph := &Adj.Hash{}
if err := graph.ReadFromFile("g.txt"); err != nil {
    panic(err)
}

// 基本操作
fmt.Println("顶点数:", graph.V())
fmt.Println("边数:", graph.E())
fmt.Println("0的邻接点:", graph.LinkedVertex(0))
fmt.Println("边(0,1)存在:", graph.HasEdge(0, 1))
fmt.Println("顶点0的度:", graph.Degree(0))
```

### 运行测试

```bash
go test ./Graph_algo/Adj
```

### 选择建议

| 场景 | 推荐实现 |
|-----|---------|
| 稠密图（边数接近 V²） | Matrix |
| 稀疏图（边数远小于 V²） | Hash |
| 需要频繁查询边 | Hash |
| 内存受限 | Table 或 Hash |
