package main

import (
	"fmt"
	"github.com/imanner/gin-xinshulaila/pkg/setting"
	"github.com/imanner/gin-xinshulaila/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
	返回Gin的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	*/
	//router := gin.Default()
	router := routers.InitRouter()
	/*
	创建不同的HTTP方法绑定到Handlers中，也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的Restful方法
	*/
	router.GET("/hello", func(c *gin.Context) {
		/*
		Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等，
		在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
		*/
		c.JSON(200, gin.H{
			"message": "go-gin",
		})
	})

	// 配置服务器
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,  // 路由处理器
		ReadTimeout:    setting.ReadTimeout,    // 读超时
		WriteTimeout:   setting.WriteTimeout,   // 写超时
		MaxHeaderBytes: 1 << 20,  // 头信息最大长度
	}

	s.ListenAndServe()
}