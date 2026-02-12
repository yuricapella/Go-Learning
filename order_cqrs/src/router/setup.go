package router

import (
	"github.com/gin-gonic/gin"
	customerRoutes "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/router"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configure trusted proxies for security
	// In development: trust no proxies (nil)
	// In production behind reverse proxy: set specific proxy IPs
	router.SetTrustedProxies(nil)

	Configure(router)

	return router
}

func Configure(engine *gin.Engine) {
	routes := customerRoutes.Routes

	for _, route := range routes {
		handler := middlewares.Logger(route.HandlerFunction)

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
