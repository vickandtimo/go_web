package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的 query string 参数
		name := c.Query("query") //通过Query获取请求中携带的query string参数
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	r.Run(":9090")
}

//11 10min
