package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/service/user"
)

func Api(router *gin.Engine) {
	router.GET("/ping", user.Login)
}
