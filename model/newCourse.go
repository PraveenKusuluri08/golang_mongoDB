package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewCourse struct {
	Title           string             `json:"title"`
	CourseName      string             `json:"coursename"`
	CourseTags      []string           `json:"coursetags"`
	CourseArticels  []string           `json:"coursearticles"`
	CreaterInfo     []string           `json:"creatorinfo"`
	CreatorEmail    string             `json:"creatoremail"`
	CourseCreatedBy string             `json:"courseCreatedBy"`
	UserId          primitive.ObjectID `json:"userid"`
	CreatedAt       string             `json:"createdat"`
}
