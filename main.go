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
	// gin.H 是一个 map[string]interface()
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title":"users/index",
	})
}

func helloWorld(c *gin.Context){
	// 模板渲染
	// 如果不向模板里传入数据，则可以将 gin.H 设为nil
	// 由于 hello 是唯一的，所以不需要用 define
	c.HTML(http.StatusOK, "hello.html", nil)
}

func returnJson(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"GO ! HEZHO!!",
	})
}

// 使用 REST 模式来写
func main() {
	r := gin.Default()
	// 在解析模板之前，要先加载静态文件
	// 这里的静态文件包括：css, js, 图片
	r.Static("/statics", "./statics")  // 每一个 / 和 . 都不能少
	// 模板解析
	r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl", "templates/posts/hello.html")
	r.GET("/posts/index", sayHello)
	r.GET("/users/index", sayHi)
	r.GET("/posts/hello", helloWorld)
	r.GET("/json", returnJson)

	r.Run(":9000")
}

