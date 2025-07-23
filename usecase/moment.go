package usecase

import "framework/domain"

type MomentUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewMomentUsecase() *MomentUsecase {
	return &MomentUsecase{}
}

// 获取动态列表
func (u *MomentUsecase) List(req domain.MomentListRequest) (domain.MomentListResponse, error) {
	// TODO: 实现获取动态列表逻辑
	return domain.MomentListResponse{}, nil
}

// 获取动态详情
func (u *MomentUsecase) Detail(req domain.MomentDetailRequest) (domain.MomentDetail, error) {
	// TODO: 实现获取动态详情逻辑
	return domain.MomentDetail{}, nil
}

// 发布动态
func (u *MomentUsecase) Post(req domain.MomentPostRequest) error {
	// TODO: 实现发布动态逻辑
	return nil
}

// 发布评论
func (u *MomentUsecase) PostComment(req domain.CommentPostRequest) error {
	// TODO: 实现发布评论逻辑
	return nil
}
