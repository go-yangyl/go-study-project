package gin

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: NewRouter()}
}

// 添加路由 method+"-"+pattern
func (e *Engine) AddRouter(method, pattern string, handler HandlerFunc) {
	e.router.AddRouter(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.AddRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRouter("POST", pattern, handler)
}

func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.AddRouter("PUT", pattern, handler)
}

func (e *Engine) Run(addr string) error {
	log.Printf("server is running %s", addr)
	return http.ListenAndServe(addr, e)
}

// 实现http.Handler接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := NewContext(w, req)

	e.router.Handler(context)
}
