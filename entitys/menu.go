package entitys

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Id   int    `gorm:"primarykey"`
	Name string `gorm:"comment:菜单名称"`
	// Category  []MenuCategory `gorm:"many2many:menu_category_rel"`
	ImagePath []string `gorm:comment:菜单图片"`
	Status    int      `gorm:"comment:菜单状态"`
}
