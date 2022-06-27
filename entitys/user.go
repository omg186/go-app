package entitys

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"comment:用户名称"`
	Password string `gorm:"	comment:密码"`
}
