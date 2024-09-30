package models

import "time"

// Post struct defines the structure of a post
type Post struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	UserID    uint       `json:"user_id"`
}
