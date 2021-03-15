package Models

type CourseContent struct {
  ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Description string `json:"description"`
	Course_id uint `json:"coure_id"`
}

func (b *CourseContent) TableName() string {
  return "course_contents"
}
