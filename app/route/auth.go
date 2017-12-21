package route

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"devpack.cc/webb_lu/ginboilerplate/app/controller/auth"
	// "encoding/json"
	"net/http"

)

// User ...
type User struct {
	ID     string `form:"id" json:"id" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required"`
}

// LoginHandler ...
func LoginHandler(c *gin.Context) {
	var user User
	// make sure the post body format meet the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// compare id and pwd in plaintext
	// TBD: move to database and use ciphertext
	if user.ID != "webb" || user.Password != "abc" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	token, loginErr := auth.Login(user.ID)

	// if sign jwt fail
	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})

}
