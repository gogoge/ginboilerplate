package logger

import (
	"os"
	"io"
	"github.com/gin-gonic/gin"
)

//Init
func Init() {
    // Disable Console Color, you don't need console color when writing the logs to file.
    gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("logs/gin.log")
    // gin.DefaultWriter = io.MultiWriter(f)

    // Use the following code if you need to write the logs to file and console at the same time.
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)	
}