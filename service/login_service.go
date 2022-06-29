package service

import (
	"errors"
	"order-food-service/dao"
	"order-food-service/pkg/error_code"
	"order-food-service/pkg/utils"
)

type LoginService struct{}

func (s *LoginService) Login(userName string, password string) (string, error) {
	userDao := dao.UserDao{}
	user, err := userDao.GetUserNamePassword(userName, password)

	if err != nil {
		return "", errors.New(error_code.ServerError)
	}
	if user == nil {
		return "", errors.New(error_code.LoginIsInvalid)
	}
	token, tokenErr := utils.GenToken(int(user.ID), user.UserName, user.Password)
	if tokenErr != nil {
		return "", tokenErr
	}
	return token, nil
}
