package usecase

type UserUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) GetMe(userID string) (interface{}, error) {
	// TODO: 实现获取当前用户信息逻辑
	return nil, nil
}

func (u *UserUsecase) UpdateMe(userID, nickname, avatarURL, bio string) error {
	// TODO: 实现更新用户信息逻辑
	return nil
}
