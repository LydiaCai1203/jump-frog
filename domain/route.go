package domain

import "time"

// Route represents the routes table
type Route struct {
	ID          string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"size:64;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Days        int       `gorm:"not null" json:"days"`
	StartCity   string    `gorm:"size:32" json:"start_city"`
	EndCity     string    `gorm:"size:32" json:"end_city"`
	IsHot       bool      `gorm:"default:false" json:"is_hot"`
	CreatedBy   string    `gorm:"type:uuid" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
