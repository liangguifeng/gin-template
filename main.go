package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/routes"
)

func main() {
	router := gin.Default()

	routes.Api(router)

	err := router.Run()
	if err != nil {
		return
	}
}
