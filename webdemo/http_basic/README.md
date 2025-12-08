# http_basic：原生 net/http Todo 示例

## 概述

这是一个使用 Go 标准库 `net/http` 实现的最小可用 REST API 示例，演示了：

- 使用 `http.NewServeMux` 管理路由
- 手动 JSON 编解码与错误处理
- 通过中间件记录请求耗时
- 线程安全的数据存储

---

## 文件结构

```
http_basic/
└── main.go     # 完整的 HTTP 服务实现（~100行）
```

---

## 核心概念

### 1. 路由处理

```go
mux := http.NewServeMux()
mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleList(store, w, r)
    case http.MethodPost:
        handleCreate(store, w, r)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
})
```

**要点**：
- `http.NewServeMux()` 是标准库的路由复用器
- 一个路径可能需要处理多种 HTTP 方法，需手动 `switch`
- 不支持路径参数（如 `/todos/:id`）

### 2. 中间件模式

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)  // 调用下一个处理器
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// 使用
srv := &http.Server{
    Handler: loggingMiddleware(mux),
}
```

**原理**：
- 中间件接收一个 `http.Handler`，返回一个新的 `http.Handler`
- 通过包装实现横切关注点（日志、认证、限流等）
- 多个中间件可以嵌套：`middleware1(middleware2(handler))`

### 3. JSON 编解码

```go
// 解码请求体
var req struct {
    Text string `json:"text"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, "invalid json", http.StatusBadRequest)
    return
}

// 编码响应
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(todo)
```

### 4. 线程安全存储

```go
type todoStore struct {
    mu    sync.Mutex  // 保护并发访问
    next  int
    todos []Todo
}

func (s *todoStore) add(text string) Todo {
    s.mu.Lock()
    defer s.mu.Unlock()
    // 临界区操作
}
```

---

## API 接口

### GET /todos - 获取所有任务

**请求**：
```bash
curl http://localhost:8080/todos
```

**响应**：
```json
[
  { "id": 1, "text": "learn go", "done": false },
  { "id": 2, "text": "build project", "done": false }
]
```

### POST /todos - 创建任务

**请求**：
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"text":"learn go"}' \
  http://localhost:8080/todos
```

**响应**（201 Created）：
```json
{ "id": 1, "text": "learn go", "done": false }
```

**错误响应**（400 Bad Request）：
```
text required
```

---

## 启动服务

```bash
cd awesome-go-datastruct
go run ./webdemo/http_basic
```

服务启动后访问 http://localhost:8080/todos

---

## 与 minigin/Gin 的对比

| 特性 | net/http 原生 | minigin/Gin |
|------|--------------|-------------|
| 路由注册 | `mux.HandleFunc` | `engine.GET/POST` |
| 方法区分 | 手动 `switch` | 自动区分 |
| 路径参数 | ❌ 不支持 | ✅ `/users/:id` |
| Context | `w, r` 分离 | 统一 `*Context` |
| JSON | 手动 `json.Encode` | `c.JSON()` |
| 中间件 | 手动嵌套包装 | `Use()` 链式注册 |

---

## 学习要点

1. **理解 `http.Handler` 接口**：只有一个方法 `ServeHTTP(w, r)`
2. **掌握中间件装饰器模式**：函数接收 handler 返回 handler
3. **熟悉 JSON 编解码**：`json.NewEncoder/Decoder`
4. **注意并发安全**：使用 `sync.Mutex` 保护共享状态

---

## 下一步

学完这个模块后，建议继续学习 [redpacket](../redpacket/) 模块，了解更复杂的业务场景和统一错误处理。
