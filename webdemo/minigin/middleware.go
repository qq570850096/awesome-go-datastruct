package minigin

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// Logger 是一个简单的日志中间件，记录方法、路径和耗时。
func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		c.Next()
		log.Printf("%s %s %v", c.Method, c.Path, time.Since(start))
	}
}

// Recovery 捕获下游 panic 并返回 500，防止程序崩溃。
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				http.Error(c.Writer, "internal server error", http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// CORS 设置基础跨域响应头，并短路 OPTIONS 预检请求。
func CORS() HandlerFunc {
	return func(c *Context) {
		h := c.Writer.Header()
		h.Set("Access-Control-Allow-Origin", "*")
		h.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		h.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// RateLimit 使用简单的每秒计数方式实现限流。
func RateLimit(qps int) HandlerFunc {
	if qps <= 0 {
		qps = 1
	}
	var (
		mu    sync.Mutex
		ts    int64
		count int
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
			http.Error(c.Writer, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		count++
		mu.Unlock()
		c.Next()
	}
}

// Timeout 在给定时间内未完成请求则返回超时响应。
func Timeout(d time.Duration) HandlerFunc {
	if d <= 0 {
		return func(c *Context) {
			c.Next()
		}
	}
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
			http.Error(c.Writer, "request timeout", http.StatusGatewayTimeout)
		}
	}
}

