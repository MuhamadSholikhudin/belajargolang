package middlewares

// func AccessManagement(Username string) {
// 	var username string
// 	username = Username
// 	fmt.Println(username)

// 	//cari role id
// 	var user models.User
// 	users := models.DB.Where("username = ?", Username).First(&user)

// 	fmt.Println(users)

// }
/*
func AccessManagement(next http.Handler) http.Handler {
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

		// mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}
		fmt.Println("Ini claims ", claims)

		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		var t2 jwt.Token = token.Claims
		fmt.Println("Ini t2 ", t2)

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

*/
