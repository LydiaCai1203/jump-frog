package repo

type UserTripRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewUserTripRepo() *UserTripRepo {
	return &UserTripRepo{}
}

func (r *UserTripRepo) ListUserTrips(userID string) (interface{}, error) {
	// TODO: 实现获取我的所有行程
	return nil, nil
}

func (r *UserTripRepo) CreateUserTrip(userID, routeID, startDate string) error {
	// TODO: 实现创建新行程
	return nil
}

func (r *UserTripRepo) GetUserTrip(userID, tripID string) (interface{}, error) {
	// TODO: 实现获取单个行程详情
	return nil, nil
}
