package repo

import "framework/domain"

type MomentRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewMomentRepo() *MomentRepo {
	return &MomentRepo{}
}

// 获取动态列表
func (r *MomentRepo) List(req domain.MomentListRequest) (domain.MomentListResponse, error) {
	// TODO: 实现获取动态列表
	return domain.MomentListResponse{}, nil
}

// 获取动态详情
func (r *MomentRepo) Detail(req domain.MomentDetailRequest) (domain.MomentDetail, error) {
	// TODO: 实现获取动态详情
	return domain.MomentDetail{}, nil
}

// 发布动态
func (r *MomentRepo) Post(req domain.MomentPostRequest) error {
	// TODO: 实现发布动态
	return nil
}

// 发布评论
func (r *MomentRepo) PostComment(req domain.CommentPostRequest) error {
	// TODO: 实现发布评论
	return nil
}
