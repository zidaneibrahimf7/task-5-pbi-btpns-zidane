package models

import (
	"gorm.io/gorm"
)

// Photo adalah model untuk entitas foto
type Photo struct {
    gorm.Model
    Title    string
    Caption  string
    PhotoURL string
    UserID   uint     // Menambahkan UserID sebagai foreign key
    User     User     // Menambahkan relasi ke User
}
