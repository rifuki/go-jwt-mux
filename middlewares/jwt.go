package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rifuki/go-jwt-mux/config"
	"github.com/rifuki/go-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("token")

			if err != nil {
				if err == http.ErrNoCookie {
					response := map[string]string{"message": "token not found"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}
			}

			/** takes token value */
			tokenString := c.Value

			claims := &config.JWTClaim{}
			/** parse jwt token */
			token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
				return config.JWT_KEY, nil
			})

			if err != nil {
				v, _ := err.(*jwt.ValidationError)
				switch v.Errors {
				case jwt.ValidationErrorExpired:
					response := map[string]string{"message": "token expired"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				case jwt.ValidationErrorNotValidYet:
					response := map[string]string{"message": "token not valid yet"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				case jwt.ValidationErrorMalformed:
					response := map[string]string{"message": "token malformed"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				case jwt.ValidationErrorSignatureInvalid:
					response := map[string]string{"message": "token signature invalid"}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				default:
					response := map[string]string{"message": err.Error()}
					helper.ResponseJSON(w, http.StatusUnauthorized, response)
					return
				}
			}

			if !token.Valid {
				response := map[string]string{"message": "token invalid"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}

			next.ServeHTTP(w, r)
		},
	)
}
