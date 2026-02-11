package route_schema

import "github.com/gin-gonic/gin"

type Route struct {
	Path            string
	Method          string
	HandlerFunction gin.HandlerFunc
}
