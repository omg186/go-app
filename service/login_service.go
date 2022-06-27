package service

import (
	"app/pkg/error_code"
	"app/pkg/utils"
	"errors"
)

type LoginService struct{}

func (s *LoginService) Login(userName string, password string) (string, error) {
	userService := UserService{}
	user := userService.GetUserNamePassword(userName, password)
	if user == nil || user.ID == 0 {
		return "", errors.New(error_code.Text(error_code.LoginIsInvalid))
	}
	token, err := utils.GenToken(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}
