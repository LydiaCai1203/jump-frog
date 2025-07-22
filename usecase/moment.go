package usecase

type MomentUsecase struct {
	// TODO: 注入 repo 层依赖
}

func NewMomentUsecase() *MomentUsecase {
	return &MomentUsecase{}
}

func (u *MomentUsecase) ListMoments() (interface{}, error) {
	// TODO: 实现获取动态列表逻辑
	return nil, nil
}

func (u *MomentUsecase) CreateMoment(userID, content string, imageURLs []string, location string) error {
	// TODO: 实现发布动态逻辑
	return nil
}

func (u *MomentUsecase) GetMoment(momentID string) (interface{}, error) {
	// TODO: 实现获取单条动态详情逻辑
	return nil, nil
}

func (u *MomentUsecase) ListComments(momentID string) (interface{}, error) {
	// TODO: 实现获取动态评论列表逻辑
	return nil, nil
}

func (u *MomentUsecase) CreateComment(userID, momentID, content string) error {
	// TODO: 实现发布评论逻辑
	return nil
}
