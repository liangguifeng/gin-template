package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liangguifeng/gin-template/setup"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	router := gin.Default()

	setup.InitRouter(router)

	_ = router.Run()
}
