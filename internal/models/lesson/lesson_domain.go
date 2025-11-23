package lesson

type LessonDomain struct {
	ID          int64  `json:"id"`
	CourseID    int64  `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"` // URL do vídeo
	Order       int    `json:"order"`
}
