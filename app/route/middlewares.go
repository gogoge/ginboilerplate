package route

import (
	// "fmt"
	// "fmt"
	"github.com/gin-gonic/gin"
	"devpack.cc/webb_lu/ginboilerplate/app/controller/auth"
	// "encoding/json"
	"net/http"
)

// type User struct {
// 	ID     string `form:"id" json:"id" binding:"required"`
// 	Password     string `form:"password" json:"password" binding:"required"`
// }


// AuthN ...
func AuthN(c *gin.Context) {
	// var payload User
	// if err := c.ShouldBindJSON(&payload); err == nil {
	// 	if payload.ID == "webb" && payload.Password == "abc" {
	// 		c.Next()
	// 	} else {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 		c.Abort()
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	c.Abort()
	// }
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	} else {
		if _, err := auth.Verify(token); err == nil {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		
	}

}

// AuthZ ...
func AuthZ(c *gin.Context) {

	c.Next()
}