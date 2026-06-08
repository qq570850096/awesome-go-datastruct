package minigin

import (
	"encoding/json"
	"net/http"
)

type HandlerFunc func(*Context)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request

	Path   string
	Method string

	Params map[string]string

	index    int
	handlers []HandlerFunc
}

func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

func (c *Context) JSON(status int, v interface{}) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	_ = json.NewEncoder(c.Writer).Encode(v)
}

func (c *Context) String(status int, s string) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, _ = c.Writer.Write([]byte(s))
}

func (c *Context) BindJSON(obj interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(obj)
}

func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}
