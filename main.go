package main

import (
	"net/http"
	"task-5-pbi-btpns-zidane/controllers" 
	"task-5-pbi-btpns-zidane/middlewares"
    "task-5-pbi-btpns-zidane/router"
	// "github.com/gorilla/mux"
)

func main() {
    // Inisialisasi UserController
    userController := controllers.NewUserController()

    // Inisialisasi router
    r := router.SetupRouter()

    // Endpoint yang dilindungi oleh middleware Authenticate
    r.Handle("/protected-endpoint", middlewares.Authenticate(http.HandlerFunc(userController.ProtectedHandler))).Methods("GET")

    // Endpoint tanpa otentikasi
    r.HandleFunc("/public-endpoint", userController.PublicHandler).Methods("GET")

    // Mulai server HTTP
    http.ListenAndServe(":8080", r)
}
