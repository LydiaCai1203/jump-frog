package usecase

type AuthUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

func (u *AuthUsecase) Register(username, password, nickname string) error {
	// TODO: 实现注册逻辑
	return nil
}

func (u *AuthUsecase) Login(username, password string) (string, error) {
	// TODO: 实现登录逻辑，返回 token
	return "mock_token", nil
}
