package auth

import (
	// "fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func makeTimestamp() int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
} 

var mySigningKey = []byte("AllYourBase")
type MyCustomClaims struct {
	// User user `json:"user"`
	Group string
	jwt.StandardClaims
}

func Sign(group string) (string, error) {
	// type user struct {
	// 	id string
	// 	loginTime int64
	// 	name string
	// 	displayName string
	// 	position string
	// 	phone string
	// 	email string
	// 	scope string
	// }
	
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
	

	// Create the Claims
	claims := MyCustomClaims{
		group,
		jwt.StandardClaims{
			ExpiresAt: makeTimestamp() + 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err


	// tokenD, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return mySigningKey, nil
	// })
	
	// if claims, ok := tokenD.Claims.(*MyCustomClaims); ok && tokenD.Valid {
	// 	fmt.Printf("%v %v",claims.User, claims.StandardClaims.ExpiresAt)
	// } else {
	// 	fmt.Println(err)
	// }
	// c.Next()
}

// Login ...
func Login(user string) (string, error){
	// dummyMiddleware()	
	group := "normal"
	if user == "webb" {
		group = "admin"
	}

	token, err := Sign(group)

	// if err != nil {
	// 	fmt.Printf("%v", err)
	// } else {
	// 	// fmt.Printf("%v", token)
	// }
	return token, err
}

// Verify ...
func Verify(signedToken string) (string, error){
	tokenD, err := jwt.ParseWithClaims(signedToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == nil {
		if claims, ok := tokenD.Claims.(*MyCustomClaims); ok && tokenD.Valid {
			return claims.Group, nil
			// fmt.Printf("%v %v",claims.Group, claims.StandardClaims.ExpiresAt)
		}
	}
	return "", err
}
