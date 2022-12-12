package middlewares

import (
	middleware "github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
	// TokenLookup: "query:token",
})

// func IsAuthenticate(){
// 	middleware.JWTWithConfig(middleware.JWTConfig{
// 		SigningKey: []byte("secret"),
// 		// TokenLookup: "query:token",
// 	}
// }
