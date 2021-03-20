package Models

type CourseContent struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    uint   `json:"course_id"`
}

func (b *CourseContent) TableName() string {
	return "course_contents"
}
