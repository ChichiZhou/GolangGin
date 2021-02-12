package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "HEZHO, CHARGE !!!!",
	})
}

// 使用 REST 模式来写
func main() {
	r := gin.Default()
	r.GET("/", sayHello)
	r.Run(":9000")
}

