package service

import (
	"errors"
	"order-food-service/dao"
	"order-food-service/dto"
	"order-food-service/pkg/error_code"

	"github.com/jinzhu/copier"
)

type UserService struct{}

// 根据用户名密码获取用户信息
func (u *UserService) GetUserNamePassword(userName string, password string) (*dto.UserRes, error) {
	userDao := dao.UserDao{}
	dto := dto.UserRes{}
	entity, err := userDao.GetUserNamePassword(userName, password)
	if err != nil {
		return nil, errors.New(error_code.ServerError)
	}
	if entity == nil {
		return nil, errors.New(error_code.UserNotFound)
	}
	copier.Copy(&dto, entity)
	return &dto, nil
}
