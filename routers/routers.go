package routers

import (
	"gpt-cms/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// 静态文件路由
	r.Static("/static", "static")

	// 首页路由
	r.GET("/", controllers.Index)

	// 文章列表页路由
	r.GET("/articles", controllers.ArticleList)

	// 单篇文章详情页路由
	r.GET("/articles/:id", controllers.ArticleDetail)

	// 用户注册页面路由
	r.GET("/register", controllers.RegisterPage)
	r.POST("/register", controllers.Register)

	// 用户登录页面路由
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.Login)

	// 用户退出路由
	r.GET("/logout", controllers.Logout)

	// 后台管理页面路由
	admin := r.Group("/admin")
	admin.Use(AuthRequired())
	{
		admin.GET("/", controllers.AdminIndex)

		// 用户管理路由
		admin.GET("/users", controllers.UserList)
		admin.GET("/users/:id", controllers.UserDetail)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)

		// 文章管理路由
		admin.GET("/articles", controllers.ArticleAdminList)
		admin.GET("/articles/:id", controllers.ArticleAdminDetail)
		admin.PUT("/articles/:id", controllers.UpdateArticle)
		admin.DELETE("/articles/:id", controllers.DeleteArticle)
		admin.GET("/articles/new", controllers.NewArticle)
		admin.POST("/articles/new", controllers.CreateArticle)

		// 标签管理路由
		admin.GET("/tags", controllers.TagList)
		admin.POST("/tags", controllers.CreateTag)
		admin.PUT("/tags/:id", controllers.UpdateTag)
		admin.DELETE("/tags/:id", controllers.DeleteTag)
	}
}
