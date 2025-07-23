package domain

import "time"

// Moment represents the moments table
type Moment struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	Content   string    `gorm:"type:text" json:"content"`
	ImageURLs string    `gorm:"type:jsonb" json:"image_urls"`
	Location  string    `gorm:"size:128" json:"location"`
	Likes     int       `gorm:"default:0" json:"likes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
