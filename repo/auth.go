package repo

import (
	"errors"

	"framework/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

// Register 用户注册
func (r *AuthRepo) Register(req domain.RegisterRequest) (id string, err error) {
	user := domain.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.ID, nil
}

// Login 用户登录校验
func (r *AuthRepo) Login(req domain.LoginRequest) (token string, err error) {
	var user domain.User
	if err := r.db.Where("username = ? OR email = ? OR phone = ?", req.Username, req.Email, req.Phone).First(&user).Error; err != nil {
		return "", err
	}
	if user.Password != req.Password {
		return "", errors.New("invalid username, email or phone or password")
	}
	return user.ID, nil
}
