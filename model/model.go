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
	IsPaid                bool               `json:"is_paid"`
}

type CoursesBuyer struct {
	CoursesNames             []string `json:"coursesnames"`
	RecentlyBoughtCourseDate []string `json:"recentCourseDate"`
}

type Mail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Purpose string `json:"purpose"`
}
