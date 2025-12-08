# Web & HTTP 实战示例

## 概述

`webdemo` 收集了一组围绕 HTTP/Web 的最小可运行示例，目标是：

- 先用标准库 `net/http` 写出可用的 REST API
- 再在业务场景中练习并发与中间件（红包并发、统一错误响应）
- 自己实现一个迷你版 Gin 框架，最后与真实 Gin 进行对比学习

---

## 架构总览

```
┌─────────────────────────────────────────────────────────────────────┐
│                         webdemo 模块架构                             │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │  http_basic  │    │   minigin    │    │ gin_example  │          │
│  │  (原生实现)   │    │  (仿Gin框架)  │    │  (真实Gin)   │          │
│  └──────┬───────┘    └──────┬───────┘    └──────┬───────┘          │
│         │                   │                   │                   │
│         ▼                   ▼                   ▼                   │
│  ┌──────────────────────────────────────────────────────┐          │
│  │                    redpacket 业务层                   │          │
│  │  ┌─────────────────┐    ┌─────────────────┐         │          │
│  │  │  Pool (Mutex)   │    │ PoolV2 (Atomic) │         │          │
│  │  │  二倍均值算法    │    │   预分配+无锁    │         │          │
│  │  └─────────────────┘    └─────────────────┘         │          │
│  └──────────────────────────────────────────────────────┘          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

请求处理流程（minigin 框架）：

    HTTP Request
         │
         ▼
    ┌─────────┐
    │ Engine  │──────────────────────────────┐
    └────┬────┘                              │
         │                                   │
         ▼                                   │
    ┌─────────┐   ┌─────────┐   ┌─────────┐ │
    │ Logger  │──▶│Recovery │──▶│  CORS   │ │ 中间件链
    └────┬────┘   └────┬────┘   └────┬────┘ │
         │             │             │       │
         ▼             ▼             ▼       │
    ┌────────────────────────────────────┐  │
    │           Router (前缀树)           │◀─┘
    │  GET /users/:id  →  handler        │
    │  POST /api/*path →  handler        │
    └─────────────────┬──────────────────┘
                      │
                      ▼
               ┌─────────────┐
               │   Context   │
               │ - Params    │
               │ - JSON()    │
               │ - BindJSON()│
               └─────────────┘
```

---

## 目录结构

| 子目录 | 说明 | 端口 | 详细文档 |
|--------|------|------|----------|
| `http_basic` | 原生 `net/http` 实现的 Todo REST API | `:8080` | [README](./http_basic/README.md) |
| `redpacket`  | 并发红包池 + HTTP 统一拦截器 | `:8090` | [README](./redpacket/README.md) |
| `minigin`    | 迷你 Gin 风格框架 + 红包服务示例 | `:8092` | [README](./minigin/README.md) |
| `gin_example`| 使用 Gin 重写红包服务，便于对比 | `:8093` | [README](./gin_example/README.md) |

---

## 三种实现方式对比

### 功能对比

| 特性 | net/http 原生 | minigin | Gin |
|------|--------------|---------|-----|
| **路由注册** | `mux.HandleFunc` | `engine.GET/POST` | `router.GET/POST` |
| **参数路由** | ❌ 不支持 | ✅ `:param` `*wildcard` | ✅ 完整支持 |
| **中间件** | 手动包装 | ✅ `Use()` 链式 | ✅ `Use()` 链式 |
| **Context** | `w, r` 分离 | ✅ 统一封装 | ✅ 功能丰富 |
| **JSON 响应** | 手动编码 | `c.JSON()` | `c.JSON()` |
| **参数绑定** | 手动解析 | `c.BindJSON()` | `c.ShouldBindJSON()` |
| **Panic 恢复** | 需手动实现 | ✅ `Recovery()` | ✅ 内置 |
| **学习价值** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |
| **生产可用** | ⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

### 代码量对比

| 模块 | 文件数 | 代码行数 | 说明 |
|------|--------|----------|------|
| http_basic | 1 | ~100 | 最小可用示例 |
| minigin | 4 | ~250 | 框架核心实现 |
| gin (框架) | - | ~10000+ | 生产级框架 |

### 代码风格对比

```go
// ========== 原生 net/http ==========
mux := http.NewServeMux()
mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
})
http.ListenAndServe(":8080", mux)

// ========== minigin ==========
engine := minigin.New()
engine.Use(minigin.Logger(), minigin.Recovery())
engine.GET("/todos", func(c *minigin.Context) {
    c.JSON(200, todos)
})
http.ListenAndServe(":8092", engine)

// ========== Gin ==========
router := gin.Default()
router.GET("/todos", func(c *gin.Context) {
    c.JSON(200, todos)
})
router.Run(":8093")
```

---

## 性能测试结果

### 红包池性能对比

测试环境：Intel Core Ultra 5 125H, Windows, Go 1.21+

```bash
go test ./webdemo/redpacket -bench=. -benchmem
```

| 版本 | 操作耗时 | 内存分配 | 说明 |
|------|----------|----------|------|
| `Pool` (Mutex) | 3,898,736 ns/op | 24 B/op | 每次 Grab 加锁 |
| `PoolV2` (Atomic) | 3,236,038 ns/op | 802,863 B/op | 预分配+原子操作 |

**结论**：
- `PoolV2` 在 Grab 阶段快约 **17%**（避免了锁竞争）
- `PoolV2` 内存占用更高（预分配所有红包金额）
- 选择建议：红包数量大且并发高时用 `PoolV2`，否则用 `Pool`

### 并发性能

```
BenchmarkPoolGrabParallel-18    18    65,181,289 ns/op
```

100,000 个红包的并发抢夺约 65ms 完成，满足大多数业务场景。

---

## API 接口文档

### 红包服务 API

所有红包服务（redpacket、minigin、gin_example）共享相同的 API 设计：

#### 初始化红包池

```http
POST /redpacket/init
Content-Type: application/json

{
  "total_amount": 1000,  // 总金额（分）
  "count": 100           // 红包个数
}
```

**成功响应**：
```json
{
  "code": 0,
  "msg": "ok",
  "data": { "ok": true }
}
```

**失败响应**：
```json
{
  "code": 400,
  "msg": "totalAmount must be at least count (1 unit each)"
}
```

#### 抢红包

```http
POST /redpacket/grab
```

**成功响应**：
```json
{
  "code": 0,
  "msg": "ok",
  "data": { "amount": 15 }  // 抢到的金额（分）
}
```

**失败响应（红包已抢完）**：
```json
{
  "code": 400,
  "msg": "no red packets left"
}
```

### Todo API（http_basic）

#### 获取所有 Todo

```http
GET /todos
```

**响应**：
```json
[
  { "id": 1, "text": "learn go", "done": false },
  { "id": 2, "text": "build project", "done": true }
]
```

#### 创建 Todo

```http
POST /todos
Content-Type: application/json

{
  "text": "learn go"
}
```

**响应**：
```json
{ "id": 1, "text": "learn go", "done": false }
```

---

## 快速启动

```bash
# 进入仓库根目录
cd awesome-go-datastruct

# 启动原生 HTTP 服务
go run ./webdemo/http_basic
# 访问 http://localhost:8080/todos

# 启动原生红包服务
go run ./webdemo/redpacket/cmd/server
# 访问 http://localhost:8090/redpacket/init

# 启动 minigin 红包服务
go run ./webdemo/minigin/cmd/redpacket
# 访问 http://localhost:8092/redpacket/init

# 启动 Gin 红包服务
go run ./webdemo/gin_example/redpacket
# 访问 http://localhost:8093/redpacket/init
```

### 测试命令

```bash
# 初始化红包（任选一个端口）
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":100}' \
  localhost:8090/redpacket/init

# 抢红包
curl -X POST localhost:8090/redpacket/grab

# 运行单元测试
go test ./webdemo/... -v

# 运行性能基准测试
go test ./webdemo/redpacket -bench=. -benchmem
```

---

## 推荐学习路径

```
     基础                  进阶                   实战
┌───────────┐      ┌───────────────┐      ┌─────────────┐
│http_basic │ ──▶  │   redpacket   │ ──▶  │   minigin   │
│  原生HTTP  │      │ 并发+拦截器    │      │  框架设计   │
└───────────┘      └───────────────┘      └──────┬──────┘
                                                  │
                                                  ▼
                                          ┌─────────────┐
                                          │ gin_example │
                                          │  对比学习    │
                                          └─────────────┘
```

1. **http_basic**：熟悉 `net/http` 原始形态与 handler 模型
2. **redpacket**：在业务场景中练习并发控制与统一拦截器
3. **minigin**：实现自己的路由树、Context 和中间件体系
4. **gin_example**：理解成熟框架在 API 设计与中间件上的取舍

配合 `BasicGo` 模块里的 `context`、`GoRoutine`、`errors`、`testingdemo` 一起看，效果更好。

---

## 相关模块

| 模块 | 关联知识点 |
|------|-----------|
| [BasicGo/context](../BasicGo/context/) | context.WithTimeout 超时控制 |
| [BasicGo/GoRoutine](../BasicGo/GoRoutine/) | sync.Mutex、WaitGroup |
| [BasicGo/errors](../BasicGo/errors/) | 错误包装与处理 |
| [BasicGo/channelselect](../BasicGo/channelselect/) | select 多路复用 |
