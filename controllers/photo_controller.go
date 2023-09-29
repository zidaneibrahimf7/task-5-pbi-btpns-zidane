package controllers

import (
	"encoding/json"
	"net/http"
	"task-5-pbi-btpns-zidane/database"
	"task-5-pbi-btpns-zidane/models"
	"github.com/gorilla/mux"
)

// PhotoController adalah controller untuk endpoint terkait foto (photo)
type PhotoController struct{}

// CreatePhoto digunakan untuk membuat foto baru
func (pc *PhotoController) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DB.Create(&photo).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(photo)
}

// GetPhotos digunakan untuk mendapatkan daftar foto
func (pc *PhotoController) GetPhotos(w http.ResponseWriter, r *http.Request) {
	var photos []models.Photo
	err := database.DB.Find(&photos).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(photos)
}

// UpdatePhoto digunakan untuk memperbarui foto berdasarkan photoId
func (pc *PhotoController) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoID := vars["photoId"]

	var updatedPhoto models.Photo
	err := json.NewDecoder(r.Body).Decode(&updatedPhoto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingPhoto models.Photo
	err = database.DB.First(&existingPhoto, photoID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update atribut foto, saya menulis kode berikut
	existingPhoto.Title = updatedPhoto.Title
	existingPhoto.Caption = updatedPhoto.Caption
	existingPhoto.PhotoURL = updatedPhoto.PhotoURL

	err = database.DB.Save(&existingPhoto).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingPhoto)
}

// DeletePhoto digunakan untuk menghapus foto berdasarkan photoId
func (pc *PhotoController) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoID := vars["photoId"]

	var existingPhoto models.Photo
	err := database.DB.First(&existingPhoto, photoID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = database.DB.Delete(&existingPhoto).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
