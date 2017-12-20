package route

import (
	"github.com/gin-gonic/gin"
)

//AuthInit ...
func AuthInit(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auth.GET("/submit", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auth.GET("/read", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

//ArticleInit ...
func ArticleInit(router *gin.Engine) {
	auth := router.Group("/articles")
	auth.GET("/read", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auth.GET("/create", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	auth.GET("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}