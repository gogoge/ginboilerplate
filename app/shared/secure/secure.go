package secure

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"	
)

//Init
func Init(host string) gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny: true,
		ForceSTSHeader: true,
		BrowserXssFilter: true,
		// SSL
		SSLRedirect: true,
		SSLHost:     host,
		// Dev mode
		IsDevelopment: false,
	})
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	return secureFunc
}