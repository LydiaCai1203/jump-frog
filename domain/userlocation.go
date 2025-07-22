package domain

import "time"

// UserLocation represents the user_locations table
type UserLocation struct {
	ID         string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID     string    `gorm:"type:uuid;not null" json:"user_id"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	RecordedAt time.Time `gorm:"autoCreateTime" json:"recorded_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
