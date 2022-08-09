package routes

import (
	UserController "CRUD/pkg/domain/users/Controllers"
	AuthController "CRUD/pkg/domain/users/Controllers/auth"
	"github.com/gorilla/mux"
)

func Api(router *mux.Router) {
	// Api Routes initialization
	apiRoutes := router.PathPrefix("/api").Subrouter()
	// auth Routes
	apiRoutes.HandleFunc("/auth/register", AuthController.Register).Methods("POST")
	apiRoutes.HandleFunc("/auth/login", AuthController.Login).Methods("POST")
	// Users Routes
	apiRoutes.HandleFunc("/users", UserController.GetAll).Methods("GET")
	apiRoutes.HandleFunc("/users/{id}", UserController.GetById).Methods("GET")
	apiRoutes.HandleFunc("/users/{id}", UserController.Update).Methods("PUT")
	apiRoutes.HandleFunc("/users/{id}", UserController.Delete).Methods("DELETE")
}
