package course

import (
	"time"
)

type CourseDomain struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Instructor  string    `json:"instructor" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
