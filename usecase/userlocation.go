package usecase

type UserLocationUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserLocationUsecase() *UserLocationUsecase {
	return &UserLocationUsecase{}
}

func (u *UserLocationUsecase) UploadLocation(userID string, latitude, longitude float64) error {
	// TODO: 实现上传用户定位逻辑
	return nil
}
