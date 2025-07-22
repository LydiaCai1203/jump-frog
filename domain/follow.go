package domain

import "time"

// Follow represents the follows table
type Follow struct {
	ID         string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FollowerID string    `gorm:"type:uuid;not null" json:"follower_id"`
	FolloweeID string    `gorm:"type:uuid;not null" json:"followee_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
