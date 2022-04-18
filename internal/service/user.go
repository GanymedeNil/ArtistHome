package service

import (
	"errors"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"
	"github.com/GanymedeNil/GoFrameworkBase/internal/model"
	"github.com/GanymedeNil/GoFrameworkBase/internal/request"
	"github.com/GanymedeNil/GoFrameworkBase/internal/response"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
}

func (u User) Create(user request.Login) (*response.UserResult, error) {
	var userResult response.UserResult
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	userModel := model.User{
		Name:     user.Username,
		Password: string(password),
	}
	err := global.DB.Where("name = ?", userResult.Name).FirstOrCreate(&userModel).Error

	return &userResult, err
}

func (u User) SingleByName(name string) *model.User {
	var user model.User
	result := global.DB.Model(&model.User{}).Where("name = ?", name).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("user not found#", zap.String("name", name))
		return nil
	}
	return &user
}

func (u User) Single(id uint) *response.UserResult {
	var user response.UserResult
	result := global.DB.Model(&model.User{}).Where("id = ?", id).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("user not found#", zap.Uint("id", id))
		return nil
	}
	return &user
}
