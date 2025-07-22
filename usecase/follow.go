package usecase

type FollowUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewFollowUsecase() *FollowUsecase {
	return &FollowUsecase{}
}

func (u *FollowUsecase) Follow(userID, followeeID string) error {
	// TODO: 实现关注逻辑
	return nil
}

func (u *FollowUsecase) Unfollow(userID, followeeID string) error {
	// TODO: 实现取消关注逻辑
	return nil
}
