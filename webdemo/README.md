## Web & HTTP 实战示例

### 概述

`webdemo` 收集了一组围绕 HTTP/web 的最小可运行示例，目标是：

- 先用标准库 `net/http` 写出可用的 REST API
- 再在业务场景中练习并发与中间件（红包并发、统一错误响应）
- 自己实现一个迷你版 Gin 框架，最后与真实 Gin 进行对比学习

目录结构：

| 子目录 | 说明 | 端口 |
|--------|------|------|
| `http_basic` | 原生 `net/http` 实现的 Todo REST API | `:8080` |
| `redpacket`  | 并发红包池 + HTTP 统一拦截器（原生 `net/http`） | `:8090` |
| `minigin`    | 迷你 Gin 风格框架 + 红包服务示例 | `:8092` |
| `gin_example`| 使用 Gin 重写红包服务，便于对比框架设计 | `:8093` |

> 进入仓库根目录 `awesome-go-datastruct` 后，再运行下面的命令。

---

### http_basic：原生 net/http Todo 示例

**文件**：`webdemo/http_basic/main.go`

**核心点**：
- 使用 `http.NewServeMux` 管理路由
- 手写 JSON 编解码与错误处理
- 通过中间件 `loggingMiddleware` 记录请求耗时

```bash
go run ./webdemo/http_basic

# 新建 todo
curl -X POST -H "Content-Type: application/json" \
  -d '{"text":"learn go"}' \
  localhost:8080/todos

# 查看 todo 列表
curl localhost:8080/todos
```

---

### redpacket：并发抢红包 + HTTP 统一拦截器

**核心文件**：
- 业务逻辑：`webdemo/redpacket/pool.go`
  - `Pool`：加锁版本红包池，按「二倍均值」随机派发
  - `PoolV2`：预分配金额 + `atomic` 分发，避免在 Grab 阶段加锁
- 单测与基准：`webdemo/redpacket/pool_test.go`
  - `TestPoolGrabConcurrent` / `TestPoolV2GrabConcurrent`
  - `BenchmarkPoolGrab` / `BenchmarkPoolV2Grab` / `BenchmarkPoolGrabParallel`
- HTTP 服务：`webdemo/redpacket/cmd/server/main.go`
  - 统一响应结构：`{"code":0,"msg":"ok","data":...}`
  - 中间件 `wrap`：统一 panic 恢复与错误转 JSON

**启动服务**：
```bash
go run ./webdemo/redpacket/cmd/server

# 初始化红包（总额 1000 分，100 个）
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":100}' \
  localhost:8090/redpacket/init

# 抢红包
curl -X POST localhost:8090/redpacket/grab
```

**运行测试与基准**：
```bash
go test ./webdemo/redpacket -run .
go test ./webdemo/redpacket -bench .
```

---

### minigin：迷你 Gin 风格框架

**核心文件**：
- 上下文：`webdemo/minigin/context.go`
  - `Context` 封装 `ResponseWriter` 和 `*Request`
  - 方法：`JSON`、`String`、`BindJSON`、`Next`、`Param`
- 路由与参数：`webdemo/minigin/router.go`
  - 前缀树 `node` + `router`
  - 支持 `:param` 和 `*wildcard` 路由模式
- 引擎：`webdemo/minigin/engine.go`
  - `Engine` 实现 `http.Handler`
  - 支持全局中间件 `Use`，`GET` / `POST` 路由注册
- 中间件：`webdemo/minigin/middleware.go`
  - `Logger`：记录方法、路径、耗时
  - `Recovery`：捕获 panic，返回 500
  - `CORS`：设置基础跨域头，短路 OPTIONS 预检
  - `RateLimit(qps int)`：简单每秒计数限流
  - `Timeout(d time.Duration)`：示意级请求超时控制
- 示例与测试：`webdemo/minigin/minigin_test.go`

**红包服务示例**：`webdemo/minigin/cmd/redpacket/main.go`
- 复用 `redpacket.Pool`，用 mini Gin 写同一套接口：
  - `POST /redpacket/init`
  - `POST /redpacket/grab`

**启动服务**：
```bash
go run ./webdemo/minigin/cmd/redpacket
# 监听 :8092，接口与 redpacket/cmd/server 兼容
```

**运行测试**：
```bash
go test ./webdemo/minigin -run .
```

---

### gin_example：Gin 版红包服务对比

**文件**：`webdemo/gin_example/redpacket/main.go`

**要点**：
- 使用 `gin.Default()` 启动带日志与 panic 恢复的 Engine
- 路由与 JSON 结构与 mini Gin 版本保持一致，便于对比：
  - `POST /redpacket/init`
  - `POST /redpacket/grab`
- 通过 `ShouldBindJSON` 简化参数解析

**启动服务**：
```bash
go run ./webdemo/gin_example/redpacket
# 监听 :8093
```

---

### 推荐学习路径

1. 先读 `http_basic`：熟悉 `net/http` 原始形态与 handler 模型
2. 再看 `redpacket`：在业务场景中练习并发控制与统一拦截器
3. 阅读 `minigin`：实现自己的路由树、Context 和中间件体系
4. 最后对照 `gin_example`：理解成熟框架在 API 设计与中间件上的取舍

配合 `BasicGo` 模块里的 `context`、`GoRoutine`、`errors`、`testingdemo` 一起看，效果更好。 

