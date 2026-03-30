package Restaurant

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Address string `gorm:"type:varchar(255);not null" json:"address"`
	Menu    string `gorm:"type:varchar(255);not null" json:"menu"`

	Chef Chef `json:"chef"`

	Pizzas []Pizza `gorm:"foreignKey:RestaurantID" json:"pizzas,omitempty"`

	Reviews []Review `gorm:"foreignKey:RestaurantID" json:"reviews,omitempty"`
}
