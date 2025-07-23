package repo

import (
	"encoding/json"
	"time"

	"framework/domain"

	"gorm.io/gorm"
)

type MomentRepo struct {
	db *gorm.DB
}

func NewMomentRepo(db *gorm.DB) *MomentRepo {
	return &MomentRepo{db: db}
}

// MomentList 获取动态列表
func (r *MomentRepo) MomentList(userID string, page int, limit int) (int64, []*domain.Moment, error) {
	var total int64
	err := r.db.Model(&domain.Moment{}).
		Where("user_id = ?", userID).
		Count(&total).Error
	if err != nil {
		return 0, nil, err
	}
	var moments []*domain.Moment
	err = r.db.Where("user_id = ?", userID).
		Order("created_at desc").
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&moments).Error
	if err != nil {
		return 0, nil, err
	}
	return total, moments, nil
}

// MomentDetail 获取动态详情
func (r *MomentRepo) MomentDetail(userID string, momentID string) (*domain.MomentDetail, error) {
	// 获取用户信息
	var user *domain.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	// 获取动态信息
	var moment *domain.Moment
	err = r.db.Where("id = ?", momentID).Where("user_id = ?", userID).First(&moment).Error
	if err != nil {
		return nil, err
	}

	// image urls 解析
	imageUrls := []string{}
	if moment.ImageURLs != "" {
		err = json.Unmarshal([]byte(moment.ImageURLs), &imageUrls)
		if err != nil {
			return nil, err
		}
	}

	// 统计所有评论的数量
	var commentsCount int64
	err = r.db.Model(&domain.Comment{}).
		Where("moment_id = ?", momentID).
		Count(&commentsCount).Error
	if err != nil {
		return nil, err
	}

	// 返回数据
	return &domain.MomentDetail{
		ID: moment.ID,
		User: domain.UserSimple{
			ID:        user.ID,
			Nickname:  user.Username,
			AvatarURL: user.AvatarURL,
		},
		Content:       moment.Content,
		ImageUrls:     imageUrls,
		Location:      moment.Location,
		LikesCount:    moment.Likes,
		CommentsCount: int(commentsCount),
		CreatedAt:     moment.CreatedAt.Format(time.DateTime),
	}, nil
}

// PostMoment 发布动态
func (r *MomentRepo) PostMoment(userID string, moment string, imageUrls []string, location string) error {
	imageUrlsJSON, err := json.Marshal(imageUrls)
	if err != nil {
		return err
	}
	err = r.db.Create(&domain.Moment{
		UserID:    userID,
		Content:   moment,
		ImageURLs: string(imageUrlsJSON),
		Location:  location,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// PostComment 发布评论
func (r *MomentRepo) PostComment(userID string, momentID string, content string) error {
	err := r.db.Create(&domain.Comment{
		UserID:   userID,
		MomentID: momentID,
		Content:  content,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
