package service

import (
	"app/dao"
	"app/entitys"
)

type UserService struct{}

// GetUserNamePassword
// 根据用户名密码获取用户信息
// @param username string 用户名
// @param password string 密码
// @return User
func (u *UserService) GetUserNamePassword(userName string, password string) *entitys.User {
	userDao := dao.UserDao{}
	return userDao.GetUserNamePassword(userName, password)
}
