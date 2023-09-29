package config

import (
    "os"
)

// GetJWTSecret mengembalikan secret key JWT
func GetJWTSecret() string {
    jwtSecret := os.Getenv("JWT_SECRET")
    if jwtSecret == "" {
        jwtSecret = "default-secret"
    }
    return jwtSecret
}
