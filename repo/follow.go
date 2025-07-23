package repo

import "framework/domain"

type FollowRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewFollowRepo() *FollowRepo {
	return &FollowRepo{}
}

// 关注用户
func (r *FollowRepo) Follow(req domain.FollowRequest) error {
	// TODO: 实现关注
	return nil
}

// 取消关注
func (r *FollowRepo) Unfollow(req domain.FollowRequest) error {
	// TODO: 实现取关
	return nil
}

// 获取粉丝列表
func (r *FollowRepo) Followers(req domain.UserFollowListRequest) (domain.UserFollowListResponse, error) {
	// TODO: 实现获取粉丝列表
	return domain.UserFollowListResponse{}, nil
}

// 获取关注列表
func (r *FollowRepo) Followees(req domain.UserFollowListRequest) (domain.UserFollowListResponse, error) {
	// TODO: 实现获取关注列表
	return domain.UserFollowListResponse{}, nil
}
