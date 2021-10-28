package router

import (
	"github.com/gin-gonic/gin"
	"wvCheck/reqHandler"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	cwwj := r.Group("/wx/auto")
	cwwj.POST("/login", reqHandler.Login)

	return r
}
