package middlewares

import (
	"fmt"
	"net/http"

	"belajargolang/go-jwt-postgres/config"
	"belajargolang/go-jwt-postgres/helper"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		fmt.Println("Ini r.RequestURI ", r.RequestURI)

		fmt.Println("Ini c ", c)

		// mengambil token value
		tokenString := c.Value
		fmt.Println("Ini tokenString ", tokenString)

		claims := &config.JWTClaim{}
		fmt.Println("Ini claims ", claims)
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		fmt.Println("Ini token ", token)
		fmt.Println("Ini config.JWT_KEY ", config.JWT_KEY)
		fmt.Println("Ini token.Header ", token.Header)
		fmt.Println("Ini token.Claims ", token.Claims)
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// token invalid
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				// token expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
