package domain

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// 路线列表 data（/routes）
type RouteListData struct {
	Routes []RouteItem `json:"routes"`
}

type RouteItem struct {
	RouteID   string `json:"route_id"`
	Name      string `json:"name"`
	Days      int    `json:"days"`
	StartCity string `json:"start_city"`
	EndCity   string `json:"end_city"`
	IsHot     bool   `json:"is_hot"`
}
