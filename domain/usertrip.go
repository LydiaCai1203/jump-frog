package domain

import "time"

// UserTrip represents the user_trips table
type UserTrip struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	RouteID   string    `gorm:"type:uuid;not null" json:"route_id"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	Status    string    `gorm:"size:16;default:'ongoing'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
