package main

import (
	"github.com/gin-gonic/gin"
	"devpack.cc/ginboilerplate/app/shared/secure"
	"devpack.cc/ginboilerplate/app/shared/logger"
	"devpack.cc/ginboilerplate/app/route"
)

func main() {
	logger.Init()
	router := gin.Default()
	router.Use(secure.Init())
	route.AuthInit(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
