package domain

import "time"

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Nickname  string    `json:"nickname"`
	AvatarURL string    `json:"avatar_url"`
	Email     string    `gorm:"unique" json:"email"`
	Phone     string    `gorm:"unique" json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
