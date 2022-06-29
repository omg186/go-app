package entitys

import "gorm.io/gorm"

// 分类
type Category struct {
	gorm.Model
	CategoryName string `gorm:"comment:分类名称"`
	Sort         int    `gorm:"moment:排序"`
	Status       int    `gorm:"moment:状态（0启用，1禁用）"`
}
