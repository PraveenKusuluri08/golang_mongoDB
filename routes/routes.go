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
	router.HandleFunc("/api/user/buyCourse/{id}", controllers.UpdateUserWhenCourseBought).Methods("PUT")

	//Authentication routes

	router.HandleFunc("/api/auth/signin", controllers.SingIn).Methods("POST")
	router.HandleFunc("/api/auth/signup", controllers.SignUp).Methods("POST")

	//Get Courses By categories
	router.HandleFunc("/api/user/getallcourse", controllers.GetAllCourses).Methods("GET")
	router.HandleFunc("/api/user/getsinglecourse/{id}", controllers.GetSingleCourse).Methods("GET")

	//add course to the cart

	router.HandleFunc("/api/user/addtocart/{userId}/{id}", controllers.AddCourse).Methods("POST")

	router.HandleFunc("/api/course/getAllCoursesByCategory/{category}", controllers.GetAllCoursesWithCategories).Methods("GET")

	router.HandleFunc("/api/course/deletesinglecourse/{courseId}/{userId}/{iscourseOwner}/{role}", controllers.DeleteSingleCourse).Methods("DELETE")

	return router
}
