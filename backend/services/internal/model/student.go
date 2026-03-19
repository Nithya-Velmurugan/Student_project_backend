package model

type Student struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"`
}
