package router

import (
	"github.com/gin-gonic/gin"
	"wvCheck/middleware"
	"wvCheck/reqHandler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	cwwj := r.Group("/wx/auto")
	cwwj.POST("/login", reqHandler.Login)

	return r
}
