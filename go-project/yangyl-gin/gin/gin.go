package gin

import (
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type RouterGroup struct {
	prefix     string
	engine     *Engine // all groups share a Engine instance
	middleware []HandlerFunc
}

type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup // store all groups
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = make([]*RouterGroup, 0)
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newRouterGroup := new(RouterGroup)
	newRouterGroup.prefix = prefix
	newRouterGroup.engine = group.engine
	// 为中间件服务
	group.engine.groups = append(group.engine.groups, newRouterGroup)

	return newRouterGroup
}

// 中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middleware = append(group.middleware, middlewares...)
}

// 添加路由 method+"-"+pattern
func (group *RouterGroup) AddRouter(method, pattern string, handler HandlerFunc) {
	path := group.prefix + pattern
	log.Printf("method:%s,pattern:%s", method, path)
	group.engine.router.AddRouter(method, path, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.AddRouter("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.AddRouter("POST", pattern, handler)
}

func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	group.AddRouter("PUT", pattern, handler)
}

func (group *RouterGroup) Run(addr string) error {
	log.Printf("server is running %s", addr)
	return http.ListenAndServe(addr, group.engine)
}

// 实现http.Handler接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	middleware := make([]HandlerFunc, 0)

	for _, group := range e.groups {
		if ok := strings.HasPrefix(req.URL.Path, group.prefix); ok {
			middleware = append(middleware, group.middleware...)
		}
	}

	context := NewContext(w, req)
	context.Handels = middleware

	e.router.Handler(context)

}
