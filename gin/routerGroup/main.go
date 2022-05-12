package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/hh")
	{
		v1.POST("/hi", Login)
		v1.GET("/ha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"HELLO": "WORLD",
			})
		})
	}

	r.Run(":9090")

}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
