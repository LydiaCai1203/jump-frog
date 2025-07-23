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

// 用户信息
// openapi: UserInfo
// {"id": "", "username": "", "nickname": "", "avatar_url": "", "bio": "", "created_at": "", ...}
type UserInfo struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	AvatarURL     string `json:"avatar_url"`
	Bio           string `json:"bio"`
	CreatedAt     string `json:"created_at"`
	TravelCount   int    `json:"travel_count"`
	CountryCount  int    `json:"country_count"`
	MomentCount   int    `json:"moment_count"`
	FollowerCount int    `json:"follower_count"`
	FolloweeCount int    `json:"followee_count"`
}

// 动态详情
// openapi: MomentDetail
type MomentDetail struct {
	ID   string `json:"id"`
	User struct {
		ID        string `json:"id"`
		Nickname  string `json:"nickname"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Content       string   `json:"content"`
	ImageUrls     []string `json:"image_urls"`
	Location      string   `json:"location"`
	CreatedAt     string   `json:"created_at"`
	LikesCount    int      `json:"likes_count"`
	CommentsCount int      `json:"comments_count"`
}

// 动态列表响应体
// {"list": [MomentDetail], "total": 0, "page": 1, "limit": 10}
type MomentListResponse struct {
	List  []MomentDetail `json:"list"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

// 路线详情
// openapi: RouteDetail
type RouteDetail struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Difficulty  string `json:"difficulty"`
	Distance    string `json:"distance"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	CreatedAt   string `json:"created_at"`
}

// 路线列表响应体
// {"data": [RouteDetail], ...}
type RouteListResponse struct {
	List  []RouteDetail `json:"list"`
	Total int           `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

// 行程详情
// openapi: TripDetail
type TripNode struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Status        string `json:"status"`
	CheckinTime   string `json:"checkin_time"`
	Rating        int    `json:"rating"`
	CommentsCount int    `json:"comments_count"`
}
type TripDetail struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	RouteID   string     `json:"route_id"`
	StartDate string     `json:"start_date"`
	EndDate   string     `json:"end_date"`
	Status    string     `json:"status"`
	CreatedAt string     `json:"created_at"`
	Nodes     []TripNode `json:"nodes"`
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

// 粉丝/关注用户
// {"id": "", "nickname": "", "avatar_url": "", "bio": ""}
type UserSimple struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

// 粉丝/关注列表响应体
// {"data": [UserSimple], ...}
type UserFollowListResponse struct {
	List  []UserSimple `json:"data"`
	Total int          `json:"total"`
	Page  int          `json:"page"`
	Limit int          `json:"limit"`
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
