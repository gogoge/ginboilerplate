package main

import (
	"devpack.cc/webb_lu/ginboilerplate/app/shared/server"
	"devpack.cc/webb_lu/ginboilerplate/app/shared/logger"
)

func main() {
	logger.Init()
	server.Run()
}
