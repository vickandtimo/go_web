package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的 querystring 参数
		name := c.Query("query") //通过Query获取请求中携带的querystring参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody") //取不到就用指定的默认值
		//name, ok := c.GetQuery("query") //取到值返回（值， true），取不到值返回（""，false）
		c.JSON(http.StatusOK, gin.H{ //取不到第二个参数就返回false
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}

//11 10min
