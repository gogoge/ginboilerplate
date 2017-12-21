package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"devpack.cc/webb_lu/ginboilerplate/app/shared/secure"
	"devpack.cc/webb_lu/ginboilerplate/app/route"
	"devpack.cc/webb_lu/ginboilerplate/configs"
)

// Run ...
func Run() {
	router := gin.Default()
	
	route.AuthInit(router)

	url := fmt.Sprintf("%s:%d", configs.Host, configs.Port)
	
	// secure middleware
	router.Use(secure.Init(url))

	err := router.RunTLS(url, configs.Certification, configs.PrivateKey)
	if err != nil {
		fmt.Println("Server could not be started")
	}
}