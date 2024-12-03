package middlewares

import (
	"context"
	
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)


func UserMiddleware (next http.Handler)  http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		tokenString :=  r.Header.Get("Authorization")

		if tokenString == "" {
			
			http.Error(w,"token required", http.StatusUnauthorized)
			return
		}

		jwtToken := os.Getenv("JWT_SECRET") 

		token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtToken),nil
		})

		if err != nil || !token.Valid {
			http.Error(w,"Invalid Token",http.StatusUnauthorized)
		}

		claims,ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userId,ok := claims["userId"].(string)

		if !ok {
			http.Error(w, "User ID not found in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(),"userId",userId)
		

		next.ServeHTTP(w,r.WithContext(ctx))
 	})	

}