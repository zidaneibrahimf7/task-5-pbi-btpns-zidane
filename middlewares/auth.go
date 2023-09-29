package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// Authenticate adalah middleware untuk otentikasi JWT
func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mendapatkan token dari header Authorization
        tokenString := r.Header.Get("Authorization")

        // Validasi token
        // Menggantikan config.GetJWTSecret() dengan secret key langsung
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("my-secret-key"), nil
        })


        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Token valid, lanjutkan ke handler berikutnya
        next.ServeHTTP(w, r)
    })
}
