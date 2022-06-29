package entitys

import "gorm.io/gorm"

// 用户
type User struct {
	gorm.Model
	UserName string `gorm:"comment:用户名称"`
	Password string `gorm:"comment:密码"`
	RealName string `gorm:"comment:真实名称"`
	Avatar   string `gorm:"comment:头像"`
	Desc     string `gorm:"comment:描述"`
	HomePath string `gorm:"comment:主页"`
	Roles    []Role `gorm:"many2many:user_role_rel"`
}
