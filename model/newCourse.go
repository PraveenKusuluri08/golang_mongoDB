package model

type NewCourse struct {
	Title           string   `json:"title"`
	CourseName      string   `json:"coursename"`
	CourseTags      []string `json:"coursetags"`
	CourseArticels  []string `json:"coursearticles"`
	CreaterInfo     []string `json:"creatorinfo"`
	CreatorEmail    string   `json:"creatoremail"`
	CourseCreatedBy string   `json:"courseCreatedBy"`
	UserId          string   `json:"userid"`
	CreatedAt       string   `json:"createdat"`
	CourseExists    bool     `json:"isCourseExists"`
}
