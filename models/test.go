package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func regesterRoutes() {
	r := gin.Default() //the middle ware that logs all of the requests of HTTP

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from the other side",
		})
		r.Run(":8081")
	})
}
