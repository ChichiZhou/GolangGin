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

func returnQuery(c *gin.Context){
	name := c.Query("query")
	//name := c.DefaultQuery("query", "someone")
	//name, ok := c.GetQuery("query")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func login(c *gin.Context){
	c.HTML(http.StatusOK, "form.html", nil)
}

func processLogin(c *gin.Context){
	name := c.PostForm("username")
	password := c.PostForm("password")
	//password := c.DefaultPostForm("password", "******")
	c.JSON(http.StatusOK, gin.H{
		"username": name,
		"password": password,
	})

}

// 使用 REST 模式来写
func main() {
	r := gin.Default()
	// 在解析模板之前，要先加载静态文件
	// 这里的静态文件包括：css, js, 图片
	r.Static("/statics", "./statics")  // 每一个 / 和 . 都不能少
	// 模板解析
	r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl", "templates/gets/hello.html", "templates/posts/form.html")
	r.GET("/posts/index", sayHello)
	r.GET("/users/index", sayHi)
	r.GET("/posts/hello", helloWorld)
	r.GET("/json", returnJson)
	r.GET("/web", returnQuery)
	r.GET("/login", login)
	r.POST("/login", processLogin)

	r.Run(":9000")
}

