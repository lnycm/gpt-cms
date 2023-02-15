package main

import (
	"fmt"
	"net/http"

	"gpt-cms/config"
	"gpt-cms/db"
	"gpt-cms/middleware"
	"gpt-cms/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化应用程序
	initApp()

	// 注册路由和中间件
	r := gin.Default()
	r.Use(middleware.Logger())
	routes.RegisterRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf("%s:%d", config.Cfg.AppHost, config.Cfg.AppPort)
	fmt.Printf("Listening on %s\n", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		panic(err)
	}
}

func initApp() {
	// 加载配置
	config.LoadConfig()

	// 连接数据库
	db.ConnectDB()

	// 设置模板引擎
	// ...

	// 设置调试模式
	gin.SetMode(config.Cfg.AppEnv)
}
