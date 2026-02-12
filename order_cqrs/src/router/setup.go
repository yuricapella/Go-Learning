package router

import (
	"github.com/gin-gonic/gin"
	customerRoutes "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/router"
)

// SetupRouter initializes the Gin router with default middleware and configures all routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	Configure(router)
	return router
}

// Configure registers all domain routes with their respective handlers and middlewares
func Configure(engine *gin.Engine) {
	routes := customerRoutes.Routes

	for _, route := range routes {
		handler := route.HandlerFunction

		switch route.Method {
		case "GET":
			engine.GET(route.Path, handler)
		case "POST":
			engine.POST(route.Path, handler)
		case "PUT":
			engine.PUT(route.Path, handler)
		case "DELETE":
			engine.DELETE(route.Path, handler)
		}
	}
}
