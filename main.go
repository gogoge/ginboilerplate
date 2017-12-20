package main

import (
	"github.com/gin-gonic/gin"
	"app/route"
)

func main() {
	router := gin.Default()
	route.AuthInit(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
