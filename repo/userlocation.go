package repo

type UserLocationRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewUserLocationRepo() *UserLocationRepo {
	return &UserLocationRepo{}
}

func (r *UserLocationRepo) UploadLocation(userID string, latitude, longitude float64) error {
	// TODO: 实现上传用户定位
	return nil
}
