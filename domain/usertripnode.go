package domain

import "time"

// UserTripNode represents the user_trip_nodes table
type UserTripNode struct {
	ID          string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserTripID  string     `gorm:"type:uuid;not null" json:"user_trip_id"`
	RouteNodeID string     `gorm:"type:uuid;not null" json:"route_node_id"`
	Status      string     `gorm:"size:16;default:'pending'" json:"status"`
	CheckinTime *time.Time `json:"checkin_time"`
	Comment     string     `gorm:"type:text" json:"comment"`
	Rating      *int       `gorm:"check:rating >= 0 AND rating <= 5" json:"rating"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
