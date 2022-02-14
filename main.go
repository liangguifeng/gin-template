package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/routes"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	router := gin.Default()

	routes.Api(router)

	_ = router.Run()
}
