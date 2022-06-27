package entitys

import "gorm.io/gorm"

// gorm.Model 的定义
type Menu struct {
	gorm.Model
	Name      string         `gorm:"comment:菜单名称"`
	Category  []MenuCategory `gorm:"many2many:menu_category_rel"`
	ImagePath []string       `gorm:comment:菜单图片"`
	Status    int            `gorm:"comment:菜单状态"`
}
