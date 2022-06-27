package dao

import (
	"app/entitys"
	"app/global"
)

type UserDao struct{}

func (d *UserDao) GetUserNamePassword(userName string, password string) *entitys.User {
	var entity entitys.User
	err := DB.Where("name = ? and password = ?", userName, password).First(&entity).Error
	if err != nil {
		global.AppLog.Error(err)
		return nil
	}
	return &entity
}
