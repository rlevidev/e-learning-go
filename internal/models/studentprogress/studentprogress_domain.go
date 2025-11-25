package studentprogress

type StudentProgressDomain struct {
	UserID   string `json:"user_id" gorm:"primaryKey;type:uuid;not null"`
	LessonID int64  `json:"lesson_id" gorm:"primaryKey;not null"`
	Done     bool   `json:"done" gorm:"default:false"`
}
