package entitys

import "gorm.io/gorm"

// 角色表
type Role struct {
	gorm.Model
	RoleName string `gorm:"comment:角色名称"`
	Value    string `gorm:"comment:编码"`
	Desc     string `gorm:"comment:描述信息"`
	Users    []User `gorm:"many2many:user_role_rel"`
}
