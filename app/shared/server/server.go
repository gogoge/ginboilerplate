package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"devpack.cc/webb_lu/ginboilerplate/app/shared/secure"
	"devpack.cc/webb_lu/ginboilerplate/app/route"
	"devpack.cc/webb_lu/ginboilerplate/configs"
	"github.com/dgrijalva/jwt-go"
	"time"
	// "encoding/json"
)

func makeTimestamp() int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
} 

func dummyMiddleware(c *gin.Context) {
	type user struct {
		id string
		loginTime int64
		name string
		displayName string
		position string
		phone string
		email string
		scope string
	}
	
	// devUser := user {
	// 	"whatever",
	// 	makeTimestamp(),
	// 	"whatever",
	// 	"whatever",
	// 	"whatever",
	// 	"whatever",
	// 	"whatever",
	// 	"whatever",
	// }
	// fmt.Println(devUser)
	mySigningKey := []byte("AllYourBase")

	type MyCustomClaims struct {
		User user `json:"user"`
		// Group string
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		// "nonAdmin",
		// devUser,
		user {
			"whatever",
			makeTimestamp(),
			"whatever",
			"whatever",
			"whatever",
			"whatever",
			"whatever",
			"whatever",
		},
		jwt.StandardClaims{
			ExpiresAt: makeTimestamp() + 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("%v", ss)
		
	}
	
	tokenD, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	
	if claims, ok := tokenD.Claims.(*MyCustomClaims); ok && tokenD.Valid {
		fmt.Printf("%v %v",claims.User, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	c.Next()
}

// Run ...
func Run() {



	// token, err := jwt.createToken("foo")

	app := gin.Default()

	url := fmt.Sprintf("%s:%d", configs.Host, configs.Port)	
	// secure middleware
	app.Use(secure.Init(url))
	app.Use(dummyMiddleware)
	route.AuthInit(app)
	err := app.RunTLS(url, configs.Certification, configs.PrivateKey)
	if err != nil {
		fmt.Println("Server could not be started")
	}

}