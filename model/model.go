package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name       string             `json:"username"`
	Email      string             `json:"email"`
	Password   string             `json:"password"`
	IsLoggedIn bool               `json:"isLoggedin"`
	Role       int                `json:"role"`
	CreatedAt  string             `json:"createdAt"`
	Course     *Courses           `json:"course"`
	UserId     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

type Courses struct {
	CoursesNames   []string `json:"coursesnames"`
	CoursesIds     []string `json:"coursesIds"`
	RecentlyBought []string `json:"recentlybought"`
	Archived       []string `json:"archieved"`
}
