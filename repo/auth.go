package repo

type AuthRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (r *AuthRepo) CreateUser(username, password, nickname string) error {
	// TODO: 实现用户注册
	return nil
}

func (r *AuthRepo) CheckUser(username, password string) (bool, error) {
	// TODO: 实现用户校验
	return true, nil
}
