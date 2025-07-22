package repo

type FollowRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewFollowRepo() *FollowRepo {
	return &FollowRepo{}
}

func (r *FollowRepo) Follow(userID, followeeID string) error {
	// TODO: 实现关注
	return nil
}

func (r *FollowRepo) Unfollow(userID, followeeID string) error {
	// TODO: 实现取消关注
	return nil
}
