package database

import (
    // "gorm.io/gorm"
    "task-5-pbi-btpns-zidane/models"
)

func Migrate() {
    // Migrasi tabel User
    DB.AutoMigrate(&models.User{})

    // Migrasi tabel Photo
    DB.AutoMigrate(&models.Photo{})
}
