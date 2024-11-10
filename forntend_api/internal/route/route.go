package route

import "github.com/gin-gonic/gin"

type Route interface {
	RegisterRoutes(g *gin.Engine)
}
