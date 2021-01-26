package gin

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

// 添加路由 method+"-"+pattern
func (r *router) AddRouter(method, pattern string, handler HandlerFunc) {
	parts := ParsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = new(node)
	}
	r.roots[method].Insert(pattern, parts, 0)

	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

// 处理请求
func (r *router) Handler(c *Context) {

	node, params := r.GetRouter(c.Method, c.Path)
	if node != nil {
		key := fmt.Sprintf("%s-%s", c.Method, node.pattern)
		fmt.Println(key)
		c.Param = params
		r.handlers[key](c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(c.W, "404 NOT FOUND: %s\n", c.Req.URL)
	}
}

func (r *router) GetRouter(method, path string) (*node, map[string]string) {
	parts := ParsePattern(path)
	params := make(map[string]string, 0)
	node, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := node.Search(parts, 0)

	if n != nil {
		parts := ParsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = parts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(parts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
