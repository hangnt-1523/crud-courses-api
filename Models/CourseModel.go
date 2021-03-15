package Models

type Course struct {
  ID    uint   `json:"id" gorm:"primary_key"`
  Title string `json:"title"`
}

func (b *Course) TableName() string {
  return "courses"
}
