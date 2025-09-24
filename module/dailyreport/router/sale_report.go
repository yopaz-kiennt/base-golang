package drrouter

import (
	"github.com/gin-gonic/gin"
)

func saleReportRouter(baseRouter *gin.RouterGroup) {
	saleReportRouter := baseRouter.Group("/sale-reports")
    
	saleReportRouter.POST("/create",)
}
