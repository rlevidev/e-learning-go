package studentprogress

type StudentProgressDomain struct {
	UserID   int64 `json:"user_id"`
	LessonID int64 `json:"lesson_id"`
	Done     bool  `json:"done"`
}
