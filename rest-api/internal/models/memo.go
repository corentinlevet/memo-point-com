package models

import "time"

type Memo struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
