package main

// 带有参数的HTTP服务器：Gin框架版本

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	//router.POST("/user/:name/*action", func(c *gin.Context) {
	//	c.FullPath() == "/user/:name/*action" // true
	//})

	router.Run("0.0.0.0:7001")
}
