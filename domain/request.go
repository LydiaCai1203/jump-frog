package domain

// 注册请求
// {"username": "", "password": "", "nickname": ""}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname,omitempty"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户信息请求（如有需要可扩展）
type UserMeRequest struct {
	Nickname  string `json:"nickname,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Bio       string `json:"bio,omitempty"`
}

// 路线详情请求
type RouteDetailRequest struct {
	RouteID string `json:"route_id"`
}

// 行程创建请求
type CreateUserTripRequest struct {
	RouteID   string `json:"route_id"`
	StartDate string `json:"start_date"`
}

// 行程详情请求
type UserTripDetailRequest struct {
	TripID string `json:"trip_id"`
}

// 动态详情请求
type MomentDetailRequest struct {
	MomentID string `json:"moment_id"`
}

// 评论创建请求
type CreateCommentRequest struct {
	MomentID string `json:"moment_id"`
	Content  string `json:"content"`
}

// 关注请求
type FollowRequest struct {
	FolloweeID string `json:"followee_id"`
}

// 定位上传请求
type UserLocationRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// 上传文件请求（如有需要可扩展）
// type UploadRequest struct { ... }
