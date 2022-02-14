package main

import (
	"github.com/liangguifeng/gin-template/setup"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	// 初始化路由
	server := setup.InitRouter()

	_ = server.Run()
}
