package api

import (
	"github.com/ViniciusMartinsS/manager/internal/infrastructure/api/middleware"
	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.CheckAccessToken)

	router.HandleFunc("/auth/login", handleAuthRequest).Methods("POST")

	router.HandleFunc("/tasks", handleTaskRequest).Methods("GET")
	router.HandleFunc("/tasks", handleTaskRequest).Methods("POST")
	router.HandleFunc("/tasks/{id}", handleTaskRequest).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handleTaskRequest).Methods("DELETE")

	return router
}
