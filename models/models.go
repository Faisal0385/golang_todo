package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task        string `gorm:"type:text;not null" json:"task"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `gorm:"not null" json:"user_id"`
	Status      string `gorm:"type:varchar(50);not null;default:'new'" json:"status"`
}

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(150);not null" json:"name"`
	Phone    string `gorm:"type:varchar(100);not null;unique" json:"phone"`
	Email    string `gorm:"type:varchar(200);unique" json:"email"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Image    string `gorm:"type:text" json:"image"`
	Status   string `gorm:"type:varchar(50);not null;default:'active'" json:"status"`
}
