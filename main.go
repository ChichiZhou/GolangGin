package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context){

	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title":"posts/index",
	})
}

func sayHi(c *gin.Context){

	// 模板渲染
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title":"users/index",
	})
}

// 使用 REST 模式来写
func main() {
	r := gin.Default()
	// 模板解析
	r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl")
	r.GET("/posts/index", sayHello)
	r.GET("/users/index", sayHi)
	r.Run(":9000")
}

