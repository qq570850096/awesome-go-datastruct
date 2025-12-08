package minigin

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEngineBasicRouting(t *testing.T) {
	e := New()
	e.GET("/ping", func(c *Context) {
		c.JSON(http.StatusOK, map[string]string{"msg": "pong"})
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal body error: %v", err)
	}
	if body["msg"] != "pong" {
		t.Fatalf(`body["msg"] = %q, want "pong"`, body["msg"])
	}
}

func TestMiddlewareOrder(t *testing.T) {
	e := New()
	trace := []string{}

	e.Use(func(c *Context) {
		trace = append(trace, "mw1-before")
		c.Next()
		trace = append(trace, "mw1-after")
	})
	e.Use(func(c *Context) {
		trace = append(trace, "mw2-before")
		c.Next()
		trace = append(trace, "mw2-after")
	})
	e.GET("/test", func(c *Context) {
		trace = append(trace, "handler")
		c.String(http.StatusOK, "ok")
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	want := []string{"mw1-before", "mw2-before", "handler", "mw2-after", "mw1-after"}
	if len(trace) != len(want) {
		t.Fatalf("trace length = %d, want %d (%v)", len(trace), len(want), trace)
	}
	for i := range want {
		if trace[i] != want[i] {
			t.Fatalf("trace[%d] = %q, want %q (full=%v)", i, trace[i], want[i], trace)
		}
	}
}

func TestParamRouting(t *testing.T) {
	e := New()
	e.GET("/hello/:name", func(c *Context) {
		c.JSON(http.StatusOK, map[string]string{"name": c.Param("name")})
	})
	e.GET("/assets/*filepath", func(c *Context) {
		c.JSON(http.StatusOK, map[string]string{"path": c.Param("filepath")})
	})

	// :param
	req := httptest.NewRequest(http.MethodGet, "/hello/gopher", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal body error: %v", err)
	}
	if body["name"] != "gopher" {
		t.Fatalf("name = %q, want %q", body["name"], "gopher")
	}

	// *wildcard
	req2 := httptest.NewRequest(http.MethodGet, "/assets/css/main.css", nil)
	w2 := httptest.NewRecorder()
	e.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w2.Code, http.StatusOK)
	}
	body = map[string]string{}
	if err := json.Unmarshal(w2.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal body error: %v", err)
	}
	if body["path"] != "css/main.css" {
		t.Fatalf("path = %q, want %q", body["path"], "css/main.css")
	}
}

func TestCORSAndRateLimitAndTimeout(t *testing.T) {
	e := New()
	e.Use(CORS(), RateLimit(1), Timeout(50*time.Millisecond))

	e.GET("/data", func(c *Context) {
		c.JSON(http.StatusOK, map[string]string{"ok": "true"})
	})

	// 测试 OPTIONS 预检请求
	reqOpt := httptest.NewRequest(http.MethodOptions, "/data", nil)
	wOpt := httptest.NewRecorder()
	e.ServeHTTP(wOpt, reqOpt)
	if wOpt.Code != http.StatusNoContent {
		t.Fatalf("OPTIONS status = %d, want %d", wOpt.Code, http.StatusNoContent)
	}
	if wOpt.Header().Get("Access-Control-Allow-Origin") == "" {
		t.Fatalf("CORS headers not set")
	}

	// 测试限流：qps=1，在同一秒内连续两次请求，第二次应被拒绝
	req1 := httptest.NewRequest(http.MethodGet, "/data", nil)
	w1 := httptest.NewRecorder()
	e.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Fatalf("first request status = %d, want %d", w1.Code, http.StatusOK)
	}

	req2 := httptest.NewRequest(http.MethodGet, "/data", nil)
	w2 := httptest.NewRecorder()
	e.ServeHTTP(w2, req2)
	if w2.Code != http.StatusTooManyRequests {
		t.Fatalf("second request status = %d, want %d", w2.Code, http.StatusTooManyRequests)
	}

	// 测试超时：使用单独 Engine 注入一个会睡眠的 handler
	e2 := New()
	e2.Use(Timeout(10 * time.Millisecond))
	e2.GET("/slow", func(c *Context) {
		time.Sleep(30 * time.Millisecond)
		c.String(http.StatusOK, "ok")
	})

	reqSlow := httptest.NewRequest(http.MethodGet, "/slow", nil)
	wSlow := httptest.NewRecorder()
	e2.ServeHTTP(wSlow, reqSlow)
	if wSlow.Code != http.StatusGatewayTimeout {
		t.Fatalf("slow request status = %d, want %d", wSlow.Code, http.StatusGatewayTimeout)
	}
}
