package routes

import (
	"github.com/gin-gonic/gin"
	api "go-mall/api/v1"
	"go-mall/middleware"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() //创建一个默认的 Engine 实例（核心路由器）
	// store := cookie.NewStore([]byte("something-very-secret")) //创建一个基于 cookie 的存储引擎，用于存储会话数据
	// r.Use(middleware.Cors(), middleware.Jaeger()) //解决跨越以及追踪问题
	r.Use(middleware.Cors()) //解决跨越问题
	// r.Use(sessions.Sessions("mysession", store)) //解决session问题
	r.StaticFS("/static", http.Dir("./static")) //静态资源
	v1 := r.Group("api/v1")
	{
		// PingExample godoc
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success") //返回一个 JSON 响应，状态码为 200，内容为 "success"
		})

		// 用户操作
		v1.POST("/user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

	}

	return r
}
