package entitys

import "gorm.io/gorm"

type MenuCategory struct {
	gorm.Model
	Name string `gorm:"primaryKey"`
}

func (MenuCategory) TableName() string {
	return "menu_category"
}
