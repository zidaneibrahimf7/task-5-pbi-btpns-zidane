package models

import (
    "gorm.io/gorm"
)

// User adalah model untuk entitas pengguna (user)
type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Photos   []Photo // Menambahkan relasi ke Photo
}

// LoginData adalah struktur data yang digunakan untuk menerima data login dari pengguna
type LoginData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
