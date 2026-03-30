package Entity

import (
	"gorm.io/gorm"
)

type Chef struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null" json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
}

type Ingredient struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);unique;not null" json:"name"`
}

type Review struct {
	gorm.Model
	RestaurantID uint   `json:"restaurant_id"`
	Rating       int    `gorm:"check:rating >= 1 AND rating <= 5" json:"rating"`
	Text         string `gorm:"type:text" json:"text"`

	Restaurant Restaurant `gorm:"foreignKey:RestaurantID" json:"restaurant_info,omitempty"`
}
