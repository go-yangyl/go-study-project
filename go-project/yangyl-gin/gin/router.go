package gin

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

// 添加路由 method+"-"+pattern
func (r *router) AddRouter(method, pattern string, handler HandlerFunc) {
	key := fmt.Sprintf("%s-%s", method, pattern)
	r.handlers[key] = handler
}

// 处理请求
func (r *router) Handler(c *Context) {
	key := fmt.Sprintf("%s-%s", c.Req.Method, c.Req.URL.Path)

	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(c.W, "404 NOT FOUND: %s\n", c.Req.URL)
	}
}

// 解析请求url
func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}


