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
	router.HandleFunc("/api/deleteSingleUser/{id}", controllers.DeleteSingleUser).Methods("DELETE")
	router.HandleFunc("/api/disableUser/{id}", controllers.DisableUser).Methods("PUT")
	router.HandleFunc("/api/enableUser", controllers.EnableUser).Methods("PUT")
	router.HandleFunc("/api/getSingleUser/{id}", controllers.GetSingleUserDocument)

	//Courses routes
	router.HandleFunc(`/api/user/createnewcourse/{id}`, controllers.CreateCourse).Methods("POST")
	return router
}
