package controllers

import (
	"encoding/json"
	"net/http"
	"task-5-pbi-btpns-zidane/database"
	"task-5-pbi-btpns-zidane/models"
	"github.com/gorilla/mux"
)

// UserController adalah struktur yang digunakan untuk mengelola operasi pengguna
type UserController struct{}

// NewUserController digunakan untuk membuat instansi UserController
func NewUserController() *UserController {
	return &UserController{}
}

// ProtectedHandler digunakan untuk mengelola endpoint yang dilindungi oleh middleware Authenticate
func (uc *UserController) ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    // Jika otentikasi berhasil, maka akan menghasilkan hasil berikut
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Ini adalah halaman yang dilindungi"))
}

// PublicHandler digunakan untuk mengelola endpoint tanpa otentikasi
func (uc *UserController) PublicHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Ini adalah halaman publik"))
}


// RegisterUser digunakan untuk membuat pengguna baru
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buat pengguna baru dalam database
	err = database.DB.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengembalikan respons dengan data pengguna yang telah dibuat
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser digunakan untuk proses login pengguna
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData models.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Untuk tujuan proyek, saya membuat token JWT palsu
	token := "contoh-token-jwt"

	// Mengembalikan token JWT sebagai respons
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// UpdateUser digunakan untuk memperbarui informasi pengguna berdasarkan userID
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingUser models.User
	err = database.DB.First(&existingUser, userID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update atribut pengguna yang diperlukan
	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	existingUser.Password = updatedUser.Password

	err = database.DB.Save(&existingUser).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
}

// DeleteUser digunakan untuk menghapus pengguna berdasarkan userID
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]

	var existingUser models.User
	err := database.DB.First(&existingUser, userID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = database.DB.Delete(&existingUser).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
