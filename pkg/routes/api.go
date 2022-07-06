package routes

import (
	UserController "CRUD/pkg/domain/users/controllers"
	"github.com/gorilla/mux"
)

func Api(router *mux.Router) {
	router.HandleFunc("/api/users", UserController.Create).Methods("POST")
	router.HandleFunc("/api/users", UserController.GetAll).Methods("GET")
	router.HandleFunc("/api/users/{id}", UserController.GetById).Methods("GET")
	router.HandleFunc("/api/users/{id}", UserController.Update).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", UserController.Delete).Methods("DELETE")
}
