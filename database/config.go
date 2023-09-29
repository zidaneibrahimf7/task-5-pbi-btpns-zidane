package database

import (
    "gorm.io/driver/postgres" // Menggunakan driver PostgreSQL
    "gorm.io/gorm"
)

var (
    DB *gorm.DB
)

func Connect() {
    // Konfigurasi koneksi ke database PostgreSQL
    dsn := "user=myuser password=mypassword dbname=mydb host=localhost port=5432 sslmode=disable" // Sesuaikan dengan koneksi dan database nantinya!

    // Membuka koneksi ke database
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Gagal terhubung ke database: " + err.Error())
    }
}
