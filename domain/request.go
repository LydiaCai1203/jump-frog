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

// 动态列表请求体
// {"user_id": "", "page": 1, "limit": 10}
type MomentListRequest struct {
	UserID string `json:"user_id"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}

// 发布动态请求体
// {"content": "", "image_urls": [""], "location": ""}
type MomentPostRequest struct {
	Content   string   `json:"content"`
	ImageUrls []string `json:"image_urls"`
	Location  string   `json:"location"`
}

// 发布评论请求体
// {"user_id": "", "moment_id": "", "content": ""}
type CommentPostRequest struct {
	UserID   string `json:"user_id"`
	MomentID string `json:"moment_id"`
	Content  string `json:"content"`
}

// 用户详情请求体
// {"user_id": ""}
type UserDetailRequest struct {
	UserID string `json:"user_id"`
}

// 用户信息请求体
// {"user_id": ""}
type UserMeInfoRequest struct {
	UserID string `json:"user_id"`
}

// 用户信息更新请求体
// {"user_id": "", "nickname": "", "avatar_url": "", "bio": ""}
type UserMeUpdateRequest struct {
	UserID    string `json:"user_id"`
	Nickname  string `json:"nickname,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Bio       string `json:"bio,omitempty"`
}

// 粉丝/关注列表请求体
// {"user_id": "", "page": 1, "limit": 10}
type UserFollowListRequest struct {
	UserID string `json:"user_id"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
}

// 路线列表请求体
// {"page": 1, "limit": 10}
type RouteListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// 路线详情请求体
// {"route_id": ""}
type RouteDetailRequestV2 struct {
	RouteID string `json:"route_id"`
}

// 行程选择请求体
// {"route_id": "", "start_date": "", "end_date": ""}
type UserTripChooseRequest struct {
	RouteID   string `json:"route_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// 行程详情请求体
// {"trip_id": ""}
type UserTripDetailRequestV2 struct {
	TripID string `json:"trip_id"`
}
