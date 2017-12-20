package main

import (
	"devpack.cc/ginboilerplate/app/shared/server"
	"devpack.cc/ginboilerplate/app/shared/logger"
)

func main() {
	logger.Init()
	server.Run()
}
