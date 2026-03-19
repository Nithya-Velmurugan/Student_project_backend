package model

import (
	"time"
)

type Role string

const (
	AdminRole   Role = "Admin"
	TeacherRole Role = "Teacher"
	StudentRole Role = "Student"
)

type User struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"` // Exclude from JSON responses
	Role         Role      `gorm:"type:varchar(20);not null;default:'Student'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
