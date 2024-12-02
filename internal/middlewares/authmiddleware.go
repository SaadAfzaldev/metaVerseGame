package middlewares

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	    
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {

			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

  
	  

		jwtSecret := os.Getenv("JWT_SECRET")

		token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret),nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            	return
		}

	
  
	   
	    next.ServeHTTP(w, r)
	})
}
  
