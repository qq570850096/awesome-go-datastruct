package minigin

import "net/http"

type Engine struct {
	router     *router
	middleware []HandlerFunc
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) Use(mw ...HandlerFunc) {
	e.middleware = append(e.middleware, mw...)
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
		index:   -1,
	}

	c.handlers = append(c.handlers, e.middleware...)

	if h, params := e.router.getRoute(r.Method, r.URL.Path); h != nil {
		c.Params = params
		c.handlers = append(c.handlers, h)
	} else {

		c.handlers = append(c.handlers, func(c *Context) {
			http.NotFound(c.Writer, c.Request)
		})
	}

	c.Next()
}
