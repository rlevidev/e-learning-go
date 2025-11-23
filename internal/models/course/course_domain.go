package course

type CourseDomain struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Instructor  string `json:"instructor"`
	CreatedAt   string `json:"created_at"`
}
