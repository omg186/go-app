package dao

import (
	"order-food-service/entitys"
	"order-food-service/global"

	"gorm.io/gorm"
)

type UserDao struct{}

func (d *UserDao) GetUserNamePassword(userName string, password string) (*entitys.User, error) {
	var entity entitys.User
	// err := DB.Where("user_name = ? and password = ?", userName, password).Joins().First(&entity).Error
	// err := DB.Joins("Role").First(&entity, "user_name = ? and password = ?", userName, password).Error
	err := DB.Preload("Roles").Where("user_name = ? and password = ?", userName, password).First(&entity).Error
	// DB.Re
	// if entity.ID != 0 {
	// 	return &entity, nil
	// }
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		global.AppLog.Error(err)
		// panic(err)
		return nil, err
	}
	return &entity, nil
}
