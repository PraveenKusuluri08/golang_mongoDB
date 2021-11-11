package model

type NewCourse struct {
	Title          string   `json:"title"`
	CourseName     string   `json:"coursename"`
	CourseTags     []string `json:"coursetags"`
	CourseArticels []string `json:"coursearticles"`
	CreaterInfo    []string `json:"creatorinfo"`
	CreatorEmail   string   `json:"creatoremail"`
	CourseId       string   `json:"_id,omitempty" bson:"_id,omitempty"`
}
