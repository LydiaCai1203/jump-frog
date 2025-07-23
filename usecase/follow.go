package usecase

import "framework/domain"

type FollowUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewFollowUsecase() *FollowUsecase {
	return &FollowUsecase{}
}

// 关注用户
func (u *FollowUsecase) Follow(req domain.FollowRequest) error {
	// TODO: 实现关注逻辑
	return nil
}

// 取消关注
func (u *FollowUsecase) Unfollow(req domain.FollowRequest) error {
	// TODO: 实现取关逻辑
	return nil
}

// 获取粉丝列表
func (u *FollowUsecase) Followers(req domain.UserFollowListRequest) (domain.UserFollowListResponse, error) {
	// TODO: 实现获取粉丝列表逻辑
	return domain.UserFollowListResponse{}, nil
}

// 获取关注列表
func (u *FollowUsecase) Followees(req domain.UserFollowListRequest) (domain.UserFollowListResponse, error) {
	// TODO: 实现获取关注列表逻辑
	return domain.UserFollowListResponse{}, nil
}
