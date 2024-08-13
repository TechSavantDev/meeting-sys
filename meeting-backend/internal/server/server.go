package server

import (
	"backend/internal/config"
	"backend/internal/route"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type server struct {
	engine *gin.Engine
}

func New() Server {
	return newServer()
}

func newServer() Server {
	return &server{
		engine: gin.Default(),
	}
}

func (s *server) Run() error {
	s.setupRouter(route.RouterConfig())
	return s.engine.Run(config.Config().ServerAddr())
}

func (s *server) setupRouter(routers []route.Router) {
	for _, routeConfig := range routers {
		group := s.engine.Group(routeConfig.GroupPath())
		for _, r := range routeConfig.Routes() {
			switch r.Method() {
			case route.GET:
				group.GET(r.Path(), r.Handler())
			case route.POST:
				group.POST(r.Path(), r.Handler())
			case route.PUT:
				group.PUT(r.Path(), r.Handler())
			case route.DELETE:
				group.DELETE(r.Path(), r.Handler())
			}
		}
	}
}
