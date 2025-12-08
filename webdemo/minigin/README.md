# minigin：迷你 Gin 风格 HTTP 框架

## 概述

这是一个仿 Gin 设计的极简 HTTP 框架，实现了核心功能：

- **前缀树路由**：支持 `:param` 参数和 `*wildcard` 通配符
- **中间件链**：洋葱模型，支持前置/后置处理
- **Context 封装**：统一请求上下文，简化 API
- **常用中间件**：Logger、Recovery、CORS、RateLimit、Timeout

---

## 文件结构

```
minigin/
├── engine.go        # 引擎核心，实现 http.Handler
├── context.go       # 请求上下文封装
├── router.go        # 前缀树路由实现
├── middleware.go    # 常用中间件集合
├── minigin_test.go  # 单元测试
└── cmd/
    └── redpacket/
        └── main.go  # 红包服务示例
```

---

## 设计原理

### 1. 前缀树路由（Trie Router）

#### 为什么需要前缀树？

```
路由列表：
  GET /users
  GET /users/:id
  GET /users/:id/posts
  POST /users
  GET /files/*filepath
```

**Map 实现的问题**：

- 只能精确匹配，无法处理 `:id` 参数
- 无法支持通配符 `*filepath`

**前缀树解决方案**：

```
GET 树：
            [root]
              │
           [users]
           ╱     ╲
        [空]    [:id]
                  │
               [posts]

路径 /users/123 匹配过程：
  root → users → :id（匹配 123）→ 提取 params["id"] = "123"
```

#### 核心数据结构

```go
type node struct {
    pattern  string   // 完整路由模式，如 /users/:id
    part     string   // 当前节点的路径片段，如 :id
    children []*node  // 子节点
    isWild   bool     // 是否模糊匹配（: 或 * 开头）
}

type router struct {
    roots    map[string]*node       // 每个 HTTP 方法一棵树
    handlers map[string]HandlerFunc // 路由 → 处理函数
}
```

#### 路由匹配流程

```
输入：GET /users/123/posts

1. 解析路径 → ["users", "123", "posts"]
2. 从 GET 树的 root 开始搜索
3. 匹配 "users" → 找到 [users] 节点
4. 匹配 "123" → 找到 [:id] 节点（isWild=true）
5. 匹配 "posts" → 找到 [posts] 节点
6. 返回 pattern="/users/:id/posts"
7. 提取参数 → params["id"] = "123"
```

#### 通配符处理

```go
// :param - 匹配单个路径段
GET /users/:id    → /users/123     → params["id"] = "123"
GET /users/:id    → /users/abc     → params["id"] = "abc"

// *wildcard - 匹配剩余所有路径
GET /files/*path  → /files/a/b/c   → params["path"] = "a/b/c"
```

---

### 2. 中间件洋葱模型

#### 执行顺序

```
请求 →
┌─────────────────────────────────────────────────┐
│  Logger (before)                                │
│  ┌───────────────────────────────────────────┐  │
│  │  Recovery (before)                        │  │
│  │  ┌─────────────────────────────────────┐  │  │
│  │  │  CORS (before)                      │  │  │
│  │  │  ┌───────────────────────────────┐  │  │  │
│  │  │  │                               │  │  │  │
│  │  │  │        Handler 业务逻辑        │  │  │  │
│  │  │  │                               │  │  │  │
│  │  │  └───────────────────────────────┘  │  │  │
│  │  │  CORS (after)                       │  │  │
│  │  └─────────────────────────────────────┘  │  │
│  │  Recovery (after)                         │  │
│  └───────────────────────────────────────────┘  │
│  Logger (after) - 记录耗时                      │
└─────────────────────────────────────────────────┘
← 响应
```

#### 实现原理

```go
type Context struct {
    index    int           // 当前执行到第几个处理器
    handlers []HandlerFunc // 中间件 + 最终 handler
}

// Next 执行下一个处理器
func (c *Context) Next() {
    c.index++
    if c.index < len(c.handlers) {
        c.handlers[c.index](c)
    }
}
```

#### 中间件示例

```go
func Logger() HandlerFunc {
    return func(c *Context) {
        start := time.Now()

        c.Next()  // 执行后续处理器

        // 后置逻辑：记录耗时
        log.Printf("%s %s %v", c.Method, c.Path, time.Since(start))
    }
}

func Recovery() HandlerFunc {
    return func(c *Context) {
        defer func() {
            if err := recover(); err != nil {
                // panic 恢复
                http.Error(c.Writer, "internal error", 500)
            }
        }()
        c.Next()
    }
}
```

---

### 3. Context 设计

#### 为什么需要 Context？

**原生 net/http**：

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 需要同时操作 w 和 r
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(data)
}
```

**minigin Context**：

```go
func handler(c *minigin.Context) {
    // 统一通过 Context 操作
    c.JSON(200, data)
}
```

#### Context 结构

```go
type Context struct {
    Writer   http.ResponseWriter
    Request  *http.Request

    Path     string              // 请求路径
    Method   string              // 请求方法
    Params   map[string]string   // 路由参数

    index    int                 // 中间件索引
    handlers []HandlerFunc       // 处理器链
}

// 便捷方法
func (c *Context) JSON(status int, v interface{})
func (c *Context) String(status int, s string)
func (c *Context) BindJSON(obj interface{}) error
func (c *Context) Param(key string) string
func (c *Context) Next()
```

---

### 4. Engine 引擎

#### 职责

1. 实现 `http.Handler` 接口
2. 管理路由注册
3. 管理全局中间件
4. 分发请求

#### 核心流程

```go
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 1. 创建 Context
    c := &Context{
        Writer:  w,
        Request: r,
        Path:    r.URL.Path,
        Method:  r.Method,
        index:   -1,
    }

    // 2. 组装处理器链：中间件 + 路由 handler
    c.handlers = append(c.handlers, e.middleware...)
    if handler, params := e.router.getRoute(r.Method, r.URL.Path); handler != nil {
        c.Params = params
        c.handlers = append(c.handlers, handler)
    } else {
        c.handlers = append(c.handlers, notFoundHandler)
    }

    // 3. 启动处理器链
    c.Next()
}
```

---

## 中间件详解

### Logger - 日志中间件

```go
func Logger() HandlerFunc {
    return func(c *Context) {
        start := time.Now()
        c.Next()
        log.Printf("%s %s %v", c.Method, c.Path, time.Since(start))
    }
}
```

### Recovery - Panic 恢复

```go
func Recovery() HandlerFunc {
    return func(c *Context) {
        defer func() {
            if rec := recover(); rec != nil {
                log.Printf("panic: %v", rec)
                http.Error(c.Writer, "internal server error", 500)
            }
        }()
        c.Next()
    }
}
```

### CORS - 跨域支持

```go
func CORS() HandlerFunc {
    return func(c *Context) {
        h := c.Writer.Header()
        h.Set("Access-Control-Allow-Origin", "*")
        h.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
        h.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

        // OPTIONS 预检请求直接返回
        if c.Request.Method == http.MethodOptions {
            c.Writer.WriteHeader(http.StatusNoContent)
            return  // 不调用 c.Next()，短路后续处理
        }
        c.Next()
    }
}
```

### RateLimit - 限流

```go
func RateLimit(qps int) HandlerFunc {
    var (
        mu    sync.Mutex
        ts    int64  // 当前秒的时间戳
        count int    // 当前秒的请求计数
    )
    return func(c *Context) {
        now := time.Now().Unix()
        mu.Lock()
        if ts != now {
            ts = now
            count = 0
        }
        if count >= qps {
            mu.Unlock()
            http.Error(c.Writer, "rate limit exceeded", 429)
            return
        }
        count++
        mu.Unlock()
        c.Next()
    }
}
```

### Timeout - 超时控制

```go
func Timeout(d time.Duration) HandlerFunc {
    return func(c *Context) {
        done := make(chan struct{})
        go func() {
            c.Next()
            close(done)
        }()

        select {
        case <-done:
            return
        case <-time.After(d):
            http.Error(c.Writer, "request timeout", 504)
        }
    }
}
```

---

## 使用示例

### 基础用法

```go
engine := minigin.New()

// 注册中间件
engine.Use(minigin.Logger(), minigin.Recovery())

// 注册路由
engine.GET("/ping", func(c *minigin.Context) {
    c.JSON(200, map[string]string{"message": "pong"})
})

engine.GET("/users/:id", func(c *minigin.Context) {
    id := c.Param("id")
    c.JSON(200, map[string]string{"id": id})
})

engine.POST("/users", func(c *minigin.Context) {
    var user User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(400, map[string]string{"error": err.Error()})
        return
    }
    c.JSON(201, user)
})

http.ListenAndServe(":8092", engine)
```

### 与 Gin 对比

| minigin                       | Gin                           |
| ----------------------------- | ----------------------------- |
| `minigin.New()`             | `gin.New()`                 |
| `engine.Use(...)`           | `router.Use(...)`           |
| `engine.GET(path, handler)` | `router.GET(path, handler)` |
| `c.JSON(status, v)`         | `c.JSON(status, v)`         |
| `c.BindJSON(&obj)`          | `c.ShouldBindJSON(&obj)`    |
| `c.Param("id")`             | `c.Param("id")`             |

---

## 启动与测试

### 启动红包服务

```bash
cd awesome-go-datastruct
go run ./webdemo/minigin/cmd/redpacket
# 服务监听 :8092
```

### 运行测试

```bash
go test ./webdemo/minigin -v
```

### 测试内容

- `TestEngineBasicRouting` - 基础路由测试
- `TestMiddlewareOrder` - 中间件执行顺序测试
- `TestParamRouting` - 参数路由测试
- `TestCORSAndRateLimitAndTimeout` - 中间件功能测试

---

## 与真实 Gin 的差异

| 特性         | minigin     | Gin                   |
| ------------ | ----------- | --------------------- |
| 路由组       | ❌          | ✅`router.Group()`  |
| 参数绑定     | 仅 JSON     | Query/Form/URI/Header |
| 验证器       | ❌          | ✅ 内置 validator     |
| 渲染器       | JSON/String | JSON/XML/YAML/HTML    |
| 路由冲突检测 | ❌          | ✅                    |
| 重定向       | ❌          | ✅                    |
| 文件服务     | ❌          | ✅ StaticFile         |
| 性能         | 良好        | 极致优化              |

---

## 学习要点

1. **前缀树**：理解 Trie 数据结构在路由匹配中的应用
2. **中间件模式**：掌握洋葱模型和处理器链设计
3. **接口设计**：学习如何封装 Context 简化 API
4. **框架思维**：理解 Web 框架的核心抽象

---

## 下一步

学完这个模块后，建议：

1. 阅读 [gin_example](../gin_example/) 对比真实 Gin 的使用
2. 尝试扩展 minigin，添加路由组、更多渲染器等功能
3. 阅读 Gin 源码，了解生产级框架的设计细节
