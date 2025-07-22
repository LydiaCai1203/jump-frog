package repo

type MomentRepo struct {
	// TODO: 注入 *gorm.DB
}

func NewMomentRepo() *MomentRepo {
	return &MomentRepo{}
}

func (r *MomentRepo) ListMoments() (interface{}, error) {
	// TODO: 实现获取动态列表
	return nil, nil
}

func (r *MomentRepo) CreateMoment(userID, content string, imageURLs []string, location string) error {
	// TODO: 实现发布动态
	return nil
}

func (r *MomentRepo) GetMoment(momentID string) (interface{}, error) {
	// TODO: 实现获取单条动态详情
	return nil, nil
}

func (r *MomentRepo) ListComments(momentID string) (interface{}, error) {
	// TODO: 实现获取动态评论列表
	return nil, nil
}

func (r *MomentRepo) CreateComment(userID, momentID, content string) error {
	// TODO: 实现发布评论
	return nil
}
