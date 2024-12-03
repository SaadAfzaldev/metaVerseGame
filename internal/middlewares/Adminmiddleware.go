package middlewares

import (
	"context"
	
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)



func AdminMiddleware (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			
		}

		jwtSecret := os.Getenv("JWT_SECRET") 

		token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			
			return []byte(jwtSecret),nil
		})

		

		if err != nil || !token.Valid{
			http.Error(w,"Token Invalid",http.StatusUnauthorized)
			return
		}

		claims,ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		role,ok := claims["role"].(string)

		if !ok {
			http.Error(w, "Role Not found in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(),"role",role)

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}