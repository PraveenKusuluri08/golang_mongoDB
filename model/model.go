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
}

type CoursesBuyer struct {
	CoursesNames   []string `json:"coursesnames"`
	CoursesId      string   `json:"coursesIds"`
	RecentlyBought []string `json:"recentlybought"`
	CourseArticels []string `json:"coursearticles"`
	CreatedAt      string   `json:"createdat"`
	UserId         string   `json:"userid"`
}
