package model

type NewCourse struct {
	Author            string   `json:"author"`
	AuthorOtherCourse []string `json:"authorOtherCourse"`
	Title             string   `json:"title"`
	CourseName        string   `json:"coursename"`
	CourseTags        []string `json:"coursetags"`
	CourseArticels    []string `json:"coursearticles"`
	CreaterInfo       []string `json:"creatorinfo"`
	CreatorEmail      string   `json:"creatoremail"`
	CourseCreatedBy   string   `json:"courseCreatedBy"`
	UserId            string   `json:"userid"`
	CreatedAt         string   `json:"createdat"`
	CourseExists      bool     `json:"isCourseExists"`
	CourseBoughtUser  string   `json:"courseboughtusers"`
	Category          string   `json:"category"`
	IsCourseOwner     bool     `json:"iscourseowner"`
}
type AddToCart struct {
	Price          int    `json:"price"`
	UserId         string `json:"userid"`
	Author         string `json:"author"`
	IsCourseBought string `json:"iscoursebought"`
}
