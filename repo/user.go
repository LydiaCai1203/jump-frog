package repo

import (
	"framework/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CountUsers() int {
	var count int64
	r.db.Model(&domain.User{}).Count(&count)
	return int(count)
}

func (r *UserRepo) GetMe(userID string) (interface{}, error) {
	// TODO: 实现获取当前用户信息
	return nil, nil
}

func (r *UserRepo) UpdateMe(userID, nickname, avatarURL, bio string) error {
	// TODO: 实现更新用户信息
	return nil
}
