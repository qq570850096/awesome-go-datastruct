# redpacket：并发抢红包 + HTTP 统一拦截器

## 概述

这是一个模拟微信红包功能的并发安全实现，演示了：

- **二倍均值算法**：公平随机分配红包金额
- **两种并发策略**：Mutex 加锁 vs 预分配+原子操作
- **统一响应拦截器**：分离业务逻辑和 HTTP 响应处理
- **Panic 恢复机制**：防止服务崩溃

---

## 文件结构

```
redpacket/
├── pool.go           # 红包池核心实现（Pool + PoolV2）
├── pool_test.go      # 并发测试 + 性能基准测试
└── cmd/
    └── server/
        └── main.go   # HTTP 服务（统一拦截器模式）
```

---

## 核心设计

### 1. 二倍均值算法

```
剩余金额：M，剩余红包数：N
当前红包金额 = random(1, 2 * M/N - 1)

示例：总金额100分，10个红包
第1次：random(1, 19) → 假设得到 12
第2次：剩余88，9个，random(1, 18) → ...
```

**特点**：
- 每个红包金额在 `[1, 2*平均值-1]` 范围内随机
- 保证每个红包至少有 1 个单位
- 整体分布相对均匀，但存在一定随机性

### 2. Pool（Mutex 版本）

```go
type Pool struct {
    mu              sync.Mutex
    remainingAmount int64  // 剩余金额（分）
    remainingCount  int    // 剩余红包数
}

func (p *Pool) Grab() (int64, error) {
    p.mu.Lock()
    defer p.mu.Unlock()
    // 计算当前红包金额
    // 更新剩余状态
}
```

**适用场景**：红包数量较少，并发不高

### 3. PoolV2（预分配+原子操作）

```go
type PoolV2 struct {
    amounts []int64  // 预先计算好的所有红包金额
    index   int64    // 当前分发位置（原子操作）
}

func (p *PoolV2) Init(total int64, count int) {
    // 一次性计算所有红包金额
    p.amounts = precompute(total, count)
}

func (p *PoolV2) Grab() (int64, error) {
    i := atomic.AddInt64(&p.index, 1) - 1  // 无锁获取
    return p.amounts[i], nil
}
```

**适用场景**：红包数量大，高并发抢夺

### 4. 性能对比

| 版本 | 策略 | Grab 耗时 | 内存占用 |
|------|------|-----------|----------|
| Pool | 每次加锁 | 3.9ms | 24 B/op |
| PoolV2 | 预分配+原子 | 3.2ms | 800KB/op |

**结论**：PoolV2 快约 17%，但内存占用更高

---

## HTTP 统一拦截器

### 设计模式

```
┌─────────────────────────────────────────────────┐
│                   wrap 包装器                    │
├─────────────────────────────────────────────────┤
│  1. defer recover() → panic 恢复                 │
│  2. data, err := handler(r) → 调用业务逻辑       │
│  3. if err → 返回错误 JSON                       │
│  4. else → 返回成功 JSON                         │
└─────────────────────────────────────────────────┘
```

### 核心代码

```go
// 业务处理函数签名：只关心业务逻辑
type appHandler func(r *http.Request) (interface{}, *appError)

// 统一包装器：处理 JSON 响应和错误
func wrap(h appHandler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // panic 恢复
        defer func() {
            if rec := recover(); rec != nil {
                writeJSON(w, 500, apiResponse{Code: 500, Msg: "internal error"})
            }
        }()

        // 调用业务逻辑
        data, appErr := h(r)
        if appErr != nil {
            writeJSON(w, appErr.Code, apiResponse{Code: appErr.Code, Msg: appErr.Msg})
            return
        }
        writeJSON(w, 200, apiResponse{Code: 0, Msg: "ok", Data: data})
    })
}
```

### 优势

1. **关注点分离**：业务逻辑不需要关心 HTTP 响应格式
2. **统一响应格式**：所有接口返回 `{code, msg, data}` 结构
3. **集中错误处理**：错误和 panic 统一在包装器中处理
4. **易于扩展**：添加日志、监控等只需修改包装器

---

## API 接口

### POST /redpacket/init - 初始化红包池

**请求**：
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":1000,"count":100}' \
  http://localhost:8090/redpacket/init
```

**成功响应**：
```json
{
  "code": 0,
  "msg": "ok",
  "data": { "ok": true }
}
```

**错误响应**：
```json
{
  "code": 400,
  "msg": "totalAmount must be at least count (1 unit each)"
}
```

### POST /redpacket/grab - 抢红包

**请求**：
```bash
curl -X POST http://localhost:8090/redpacket/grab
```

**成功响应**：
```json
{
  "code": 0,
  "msg": "ok",
  "data": { "amount": 15 }
}
```

**错误响应（红包抢完）**：
```json
{
  "code": 400,
  "msg": "no red packets left"
}
```

---

## 启动与测试

### 启动服务

```bash
cd awesome-go-datastruct
go run ./webdemo/redpacket/cmd/server
# 服务监听 :8090
```

### 运行测试

```bash
# 单元测试
go test ./webdemo/redpacket -v

# 性能基准测试
go test ./webdemo/redpacket -bench=. -benchmem
```

### 并发压测

```bash
# 初始化 10000 个红包
curl -X POST -H "Content-Type: application/json" \
  -d '{"total_amount":100000,"count":10000}' \
  localhost:8090/redpacket/init

# 使用 ab 或 wrk 进行压测
ab -n 10000 -c 100 -m POST localhost:8090/redpacket/grab
```

---

## 学习要点

1. **并发安全**：理解 `sync.Mutex` 和 `sync/atomic` 的使用场景
2. **随机算法**：掌握二倍均值算法的原理和实现
3. **设计模式**：学习统一拦截器模式分离关注点
4. **性能优化**：通过预分配减少锁竞争

---

## 下一步

学完这个模块后，建议继续学习 [minigin](../minigin/) 模块，了解如何构建一个完整的 HTTP 框架。
