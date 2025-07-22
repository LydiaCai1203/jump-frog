package usecase

type UserTripUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserTripUsecase() *UserTripUsecase {
	return &UserTripUsecase{}
}

func (u *UserTripUsecase) ListUserTrips(userID string) (interface{}, error) {
	// TODO: 实现获取我的所有行程逻辑
	return nil, nil
}

func (u *UserTripUsecase) CreateUserTrip(userID, routeID, startDate string) error {
	// TODO: 实现创建新行程逻辑
	return nil
}

func (u *UserTripUsecase) GetUserTrip(userID, tripID string) (interface{}, error) {
	// TODO: 实现获取单个行程详情逻辑
	return nil, nil
}
