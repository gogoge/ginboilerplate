package route

import (
	"github.com/gin-gonic/gin"
)

// Dummy ...
func Dummy(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

//AuthInit ...
func AuthInit(router *gin.Engine) {
	authRoute := router.Group("/auth")
	authRoute.POST("/login", LoginHandler)
	authRoute.GET("/submit", Dummy)
	authRoute.GET("/read", Dummy)
}

//ArticleInit ...
func ArticleInit(router *gin.Engine) {
	articleRoute := router.Group("/articles")
	articleRoute.Use(AuthN)
	articleRoute.Use(AuthZ)
	articleRoute.GET("/read", Dummy)
	articleRoute.GET("/create", Dummy)
	articleRoute.GET("/delete", Dummy)
}