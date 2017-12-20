package main

import (
	"github.com/gin-gonic/gin"
	"app/shared/secure"
	"app/route"
)

func main() {
	router := gin.Default()
	router.Use(secure.Init())
	route.AuthInit(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
