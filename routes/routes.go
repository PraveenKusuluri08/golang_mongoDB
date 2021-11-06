package routes

import (
	"github.com/PraveenKusuluri08/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/createUser", controllers.CreateUserAccout).Methods("POST")
	router.HandleFunc("/api/getAllUsers", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/updateSingleUser/{id}", controllers.UpdateSingleUser).Methods("PUT")
	return router
}
