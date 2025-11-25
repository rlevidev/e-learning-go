package lesson

type LessonDomain struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	CourseID    int64  `json:"course_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Content     string `json:"content"` // URL do vídeo
	Order       int    `json:"order" gorm:"not null"`
}
