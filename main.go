package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"devpack.cc/ginboilerplate/app/shared/secure"
	"devpack.cc/ginboilerplate/app/shared/logger"
	"devpack.cc/ginboilerplate/app/route"
	"devpack.cc/ginboilerplate/configs"
)

func main() {
	logger.Init()
	router := gin.Default()
	
	route.AuthInit(router)

	host := "localhost"
	nonHTTPSPort := 10080
	httpsPort := 10443
	httpURL := fmt.Sprintf("%s:%d", host, nonHTTPSPort)
	secureHTTPURL := fmt.Sprintf("%s:%d", host, httpsPort)
	// secure middleware
	router.Use(secure.Init(secureHTTPURL))

	// HTTP (listen and redirect HTTP to HTTPS)
	go router.Run(httpURL)

	err := router.RunTLS(secureHTTPURL, configs.Certification, configs.PrivateKey)
	if err != nil {
		fmt.Println("Server could not be started")
	}
	// router.Run() // listen and serve on 0.0.0.0:8080
}
