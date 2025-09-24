package main

import (
	"github.com/gin-gonic/gin"
	drrouter "backend/module/dailyreport/router"
)

func main() {
	router := gin.Default()
	router = drrouter.SetupRouter(router)
	router.Run(":8080")
}
