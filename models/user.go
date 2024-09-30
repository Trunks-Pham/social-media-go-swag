package models

import "time"

// User struct defines the structure of a user
type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Username  string     `json:"username" gorm:"unique"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Posts     []Post     `json:"posts" gorm:"foreignKey:UserID"`
}
