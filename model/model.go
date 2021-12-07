package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserName              string             `json:"username"`
	Email                 string             `json:"email"`
	Password              string             `json:"password"`
	IsLoggedIn            bool               `json:"isLoggedin"`
	Role                  int                `json:"role"`
	CreatedAt             string             `json:"createdAt"`
	CoursesBought         *CoursesBuyer      `json:"course"`
	IsExists              bool               `json:"isExists"`
	UserId                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NumberOfCourseCreated int                `json:"numberOfCoursesCreated"`
}

type CoursesBuyer struct {
	CoursesNames             []string `json:"coursesnames"`
	CourseArticels           []string `json:"coursearticles"`
	RecentlyBoughtCourseDate string   `json:"recentCourseDate"`
}

type BuySingleCourse struct {
	CourseID       string   `json:"courseid"`
	CourseTags     []string `json:"coursetags"`
	CourseArticles []string `json:"courseArticles"`
	Sections       []string `json:"sections"`
	Quizzes        []string `json:"quizzes"`
	Links          []string `json:"links"`
}

type Mail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Purpose string `json:"purpose"`
}
