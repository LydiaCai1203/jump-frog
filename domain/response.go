package domain

// 通用响应结构体
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// 注册响应
type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

// 登录响应
type LoginResponse struct {
	Token string `json:"token"`
}

// 用户信息响应
type UserMeResponse struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

// 路线详情响应
type RouteDetailResponse struct {
	RouteID   string `json:"route_id"`
	Name      string `json:"name"`
	Days      int    `json:"days"`
	StartCity string `json:"start_city"`
	EndCity   string `json:"end_city"`
	IsHot     bool   `json:"is_hot"`
}

// 路线列表响应
type RouteListResponse struct {
	Routes []RouteDetailResponse `json:"routes"`
}

// 行程详情响应
type UserTripDetailResponse struct {
	TripID    string `json:"trip_id"`
	RouteID   string `json:"route_id"`
	StartDate string `json:"start_date"`
	Status    string `json:"status"`
}

// 行程列表响应
type UserTripListResponse struct {
	Trips []UserTripDetailResponse `json:"trips"`
}

// 动态详情响应
type MomentDetailResponse struct {
	MomentID string `json:"moment_id"`
	Content  string `json:"content"`
	UserID   string `json:"user_id"`
}

// 动态列表响应
type MomentListResponse struct {
	Moments []MomentDetailResponse `json:"moments"`
}

// 评论详情响应
type CommentDetailResponse struct {
	CommentID string `json:"comment_id"`
	MomentID  string `json:"moment_id"`
	Content   string `json:"content"`
	UserID    string `json:"user_id"`
}

// 评论列表响应
type CommentListResponse struct {
	Comments []CommentDetailResponse `json:"comments"`
}

// 关注响应
type FollowResponse struct {
	FolloweeID string `json:"followee_id"`
	Status     string `json:"status"`
}

// 定位上传响应
type UserLocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

// 上传文件响应
type UploadResponse struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}
