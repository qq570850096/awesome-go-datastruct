package minigin

import "net/http"

// Engine 是一个极简的 HTTP 路由与中间件引擎。
type Engine struct {
	router     *router
	middleware []HandlerFunc
}

// New 创建一个新的 Engine。
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// Use 注册全局中间件，按注册顺序执行。
func (e *Engine) Use(mw ...HandlerFunc) {
	e.middleware = append(e.middleware, mw...)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET 注册 GET 路由。
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

// POST 注册 POST 路由。
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

// ServeHTTP 实现 http.Handler，负责路由与中间件链的调度。
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		index:   -1,
	}

	// 组合中间件与路由处理器。
	c.handlers = append(c.handlers, e.middleware...)

	if h, params := e.router.getRoute(r.Method, r.URL.Path); h != nil {
		c.Params = params
		c.handlers = append(c.handlers, h)
	} else {
		// 若未匹配到路由，返回 404。
		c.handlers = append(c.handlers, func(c *Context) {
			http.NotFound(c.Writer, c.Request)
		})
	}

	c.Next()
}
