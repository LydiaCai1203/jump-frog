package domain

import "time"

// RouteNode represents the route_nodes table
type RouteNode struct {
	ID          string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RouteID     string     `gorm:"type:uuid;not null" json:"route_id"`
	Day         int        `gorm:"not null" json:"day"`
	OrderInDay  int        `gorm:"not null" json:"order_in_day"`
	Name        string     `gorm:"size:64;not null" json:"name"`
	Type        string     `gorm:"size:32" json:"type"`
	Location    string     `gorm:"size:128" json:"location"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	OpenTime    *time.Time `json:"open_time"`
	CloseTime   *time.Time `json:"close_time"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
