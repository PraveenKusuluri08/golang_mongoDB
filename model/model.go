package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name                  string             `json:"username"`
	Email                 string             `json:"email"`
	Password              string             `json:"password"`
	IsLoggedIn            bool               `json:"isLoggedin"`
	Role                  int                `json:"role"`
	CreatedAt             string             `json:"createdAt"`
	CoursesBought         *CoursesBuyer      `json:"course"`
	IsExists              bool               `json:"isExists"`
	UserId                primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NumberOfCourseCreated int                `json:"numberOfCoursesCreated"`
	Token                 string             `json:"token"`
}

type CoursesBuyer struct {
	CoursesNames             []string `json:"coursesnames"`
	CourseArticels           []string `json:"coursearticles"`
	RecentlyBoughtCourseDate string   `json:"recentCourseDate"`
}
