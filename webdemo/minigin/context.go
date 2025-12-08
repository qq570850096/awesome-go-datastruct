package minigin

import (
	"encoding/json"
	"net/http"
)

// HandlerFunc 是所有处理函数和中间件的统一签名。
type HandlerFunc func(*Context)

// Context 封装每次请求的上下文，提供常用辅助方法。
type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string

	Params map[string]string

	index    int
	handlers []HandlerFunc
}

// Next 按顺序执行剩余的中间件/处理函数。
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// JSON 写入 JSON 响应。
func (c *Context) JSON(status int, v interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	_ = json.NewEncoder(c.Writer).Encode(v)
}

// String 写入纯文本响应。
func (c *Context) String(status int, s string) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, _ = c.Writer.Write([]byte(s))
}

// BindJSON 从请求体解析 JSON 到给定结构体。
func (c *Context) BindJSON(obj interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(obj)
}

// Param 返回路径参数的值，不存在时返回空字符串。
func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}

