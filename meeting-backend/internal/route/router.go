package route

import (
	"github.com/gin-gonic/gin"
)

func RouterConfig() []Router {
	return Routers
}

type Router struct {
	groupPath string
	routes    []Router
	method    MethodType
	handler   func(c *gin.Context)
	path      string
}

func (r *Router) GroupPath() string {
	return r.groupPath
}

func (r *Router) Routes() []Router {
	return r.routes
}

func (r *Router) Method() MethodType {
	return r.method
}

func (r *Router) Handler() func(c *gin.Context) {
	return r.handler
}

func (r *Router) Path() string {
	return r.path
}
