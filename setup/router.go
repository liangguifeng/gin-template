package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/middleware"
	"github.com/liangguifeng/gin-template/server/user"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api/v1").Use(middleware.Auth())
	apiRouterWithoutRecord := router.Group("/api/v1")
	{
		// 鉴权路由组
		apiRouter.GET("login", user.LoginHandler)
	}
	{
		// 不鉴权路由组
		apiRouterWithoutRecord.GET("login", user.LoginHandler)
	}

	return router
}
