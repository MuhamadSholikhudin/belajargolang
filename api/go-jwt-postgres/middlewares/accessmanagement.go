package middlewares

import (
	"net/http"

	"belajargolang/api/go-jwt-postgres/config"
	"belajargolang/api/go-jwt-postgres/helper"
	"belajargolang/api/go-jwt-postgres/models"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AccessMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		// mengambil token value
		tokenString := c.Value
		claims := &config.JWTClaim{}
		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if !token.Valid {

		}
		var user models.User
		if err := models.DB.First(&user, "username = ?", claims.Username).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
		var role models.Role
		if err := models.DB.First(&role, "id = ?", user.Role_id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
		var access_menu models.Access_menu
		if err := models.DB.First(&access_menu, "role_id = ?", user.Role_id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
		var menu models.Menu
		if err := models.DB.First(&menu, "id = ?", access_menu.Menu_id).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
		var sub_menu models.Sub_Menu
		if err := models.DB.First(&sub_menu, "id = ? AND url = ?", menu.Id, r.RequestURI).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": err.Error()}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
