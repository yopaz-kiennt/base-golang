package drrouter

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/api/v1")

	saleReportRouter(v1)
	// Thêm các router khác ở đây

	return router
}
