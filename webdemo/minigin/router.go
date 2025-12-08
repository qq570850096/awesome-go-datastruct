package minigin

import (
	"strings"
)

// node 是路由前缀树上的节点。
type node struct {
	pattern  string   // 完整路由，例如 /p/:lang
	part     string   // 路径片段，例如 :lang
	children []*node  // 子节点
	isWild   bool     // 是否模糊匹配，: 或 *
}

// matchChild 返回第一个匹配的子节点（用于插入）。
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 返回所有匹配的子节点（用于搜索）。
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 将一个路由模式插入前缀树。
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*"),
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search 在前缀树中查找匹配的节点。
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		if result := child.search(parts, height+1); result != nil {
			return result
		}
	}
	return nil
}

// router 管理不同 HTTP 方法的前缀树以及最终 handler。
type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern 将路由模式按 / 拆分，并处理 * 通配符。
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item == "" {
			continue
		}
		parts = append(parts, item)
		if strings.HasPrefix(item, "*") {
			break
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRoute 匹配路由并提取路径参数。
func (r *router) getRoute(method, path string) (HandlerFunc, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	searchParts := parsePattern(path)
	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	patternParts := parsePattern(n.pattern)
	params := make(map[string]string)
	for idx, part := range patternParts {
		if strings.HasPrefix(part, ":") {
			params[part[1:]] = searchParts[idx]
		}
		if strings.HasPrefix(part, "*") {
			params[part[1:]] = strings.Join(searchParts[idx:], "/")
			break
		}
	}

	key := method + "-" + n.pattern
	return r.handlers[key], params
}

