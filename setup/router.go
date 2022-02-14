package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/server/user"
)

func InitRouter(router *gin.Engine) {
	// todo lianggf 中间件
	router.GET("/ping", user.Login)
}
