package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lokesh-go/go-boilerplate/api/rest/handler"
	"github.com/lokesh-go/go-boilerplate/api/rest/middleware"
	"github.com/lokesh-go/go-boilerplate/internal/dependencies"
)

// Router interface
type Router interface {
	RegisterPublicRoutes() *gin.Engine
	RegisterInternalRoutes() *gin.Engine
}

type router struct {
	handler *handler.Handler
}

// New router
func New(d *dependencies.Deps) Router {
	return &router{
		handler: handler.New(d),
	}
}

// Register public routes
func (r *router) RegisterPublicRoutes() *gin.Engine {
	// Initialise gin router
	router := gin.New()

	// Use middlewares
	router.Use(middleware.CustomLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.CaptureRequestMetaData())

	// Define routes
	router.GET("/ping", r.handler.Ping)

	/*
		public := router.Group("/api/v1")
		{
			user := public.Group("/users")
			{
				user.GET("/:userId", r.handler.GetUser)
			}
		}
	*/

	// Returns
	return router
}

// Register internal routes
func (r *router) RegisterInternalRoutes() *gin.Engine {
	// Initialise gin router
	router := gin.New()

	// Use middlewares
	router.Use(middleware.CustomLogger())
	router.Use(gin.Recovery())
	router.Use(middleware.CaptureRequestMetaData())

	// Define routes
	router.GET("/ping", r.handler.Ping)

	/*
		public := router.Group("/internal/api/v1")
		{
		}
	*/

	// Returns
	return router
}
