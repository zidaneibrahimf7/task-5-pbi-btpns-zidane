package router

import (
    "github.com/gorilla/mux"
    "net/http"
    "task-5-pbi-btpns-zidane/middlewares"
    "task-5-pbi-btpns-zidane/controllers"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Endpoint yang dilindungi oleh middleware Authenticate
    r.Handle("/protected-endpoint", middlewares.Authenticate(http.HandlerFunc(controllers.NewUserController().ProtectedHandler))).Methods("GET")

    // Endpoint tanpa otentikasi
    r.HandleFunc("/public-endpoint", controllers.NewUserController().PublicHandler).Methods("GET")

    // Gabungkan endpoint dari controller Anda
    uc := &controllers.UserController{}
    pc := &controllers.PhotoController{}

    // Endpoint untuk pengguna (user)
    r.HandleFunc("/users/register", uc.RegisterUser).Methods("POST")
    r.HandleFunc("/users/login", uc.LoginUser).Methods("POST")
    r.HandleFunc("/users/{userId}", uc.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{userId}", uc.DeleteUser).Methods("DELETE")

    // Endpoint untuk foto (photo)
    r.HandleFunc("/photos", pc.CreatePhoto).Methods("POST")
    r.HandleFunc("/photos", pc.GetPhotos).Methods("GET")
    r.HandleFunc("/photos/{photoId}", pc.UpdatePhoto).Methods("PUT")
    r.HandleFunc("/photos/{photoId}", pc.DeletePhoto).Methods("DELETE")

    return r
}
