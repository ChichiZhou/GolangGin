package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct{
	Username string `form:"name"`
	Age string `form:"age"`
}
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

func middleWare(c *gin.Context){
	fmt.Println("You are in middleWare NO.1")
}

func handleIndex(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Everything is OK",
	})
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
	name := c.PostForm("name")
	password := c.PostForm("password")
	//password := c.DefaultPostForm("password", "******")
	c.JSON(http.StatusOK, gin.H{
		"username": name,
		"password": password,
	})
}

func processParam(c *gin.Context){
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age": age,
	})
}

func paramBind(c *gin.Context){
	var u UserInfo
	c.ShouldBind(&u)
	fmt.Println(u)
	c.JSON(http.StatusOK, gin.H{
		"name": u.Username,
		"age":u.Age,
	})

}

// 使用 REST 模式来写
func main() {
	// 生成默认的路由
	r := gin.Default()
	// 在解析模板之前，要先加载静态文件
	// 这里的静态文件包括：css, js, 图片
	r.Static("/statics", "./statics")  // 每一个 / 和 . 都不能少
	// 模板解析
	r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl", "templates/gets/hello.html", "templates/posts/form.html","templates/views/404.html")
	// 设置路由
	r.GET("/posts/index", sayHello)
	r.GET("/users/index", sayHi)
	r.GET("/posts/hello", helloWorld)
	r.GET("/json", returnJson)
	r.GET("/web", returnQuery)

	r.GET("/login", login)
	r.POST("/login", processLogin)

	// 这里的 : 相当于通配符，所以/:name/:age 实际上和上面的 /posts/index 冲突了
	// 所以这里要加上一个 /user
	r.GET("/user/:name/:age", processParam)
	// 注意这里输入的是传入参数
	r.GET("/user", paramBind)

	r.GET("/middleware", middleWare, handleIndex)
	r.NoRoute(func(c *gin.Context){
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run(":9000")
}

