package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会走这一条处理逻辑
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//Any: 请求方法的大集合/大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			//...
		}
	})

	//NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "CBiBank.com"})
	})

	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})

	//路由组 多用于区分不同的业务线或API版本
	//把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}

	//商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index "})
	})
	r.Run(":9090")
}
